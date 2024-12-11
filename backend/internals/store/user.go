package store

import (
	"backend/internals/models"
	"context"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserStore interface {
	ChangePassword(ctx context.Context, userID int, currentPassword string, newPassword string) error
	CreateUser(ctx context.Context, user *models.User) error
	DeleteUserByID(ctx context.Context, userID int) (string, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	GetUserByID(ctx context.Context, id int) (*models.User, error)
	GetUserByUserName(ctx context.Context, username string) (*models.User, error)
	UpdateProfilePicture(ctx context.Context, id int, profilePicture string) error
	UpdateUserPassword(ctx context.Context, user *models.User) error
	UpdateUserProfile(ctx context.Context, user *models.User) (*models.User, error)
}

type SQLiteUserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *SQLiteUserStore {
	return &SQLiteUserStore{db: db}
}

func (s *SQLiteUserStore) CreateUser(ctx context.Context, user *models.User) error {
	_, err := s.db.ExecContext(ctx, `INSERT INTO users (username, password, email) VALUES (?, ?, ?)`, user.Username, user.Password, user.Email)
	return err
}

func (s *SQLiteUserStore) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	var profilePicture sql.NullString // Use sql.NullString to handle NULL values
	err := s.db.QueryRowContext(ctx, "SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.IsAdmin, &profilePicture)
	if err != nil {
		return nil, err
	}
	if profilePicture.Valid {
		user.ProfilePicture = &profilePicture.String
	} else {
		user.ProfilePicture = nil
	}
	return &user, nil
}

func (s *SQLiteUserStore) GetUserByUserName(ctx context.Context, username string) (*models.User, error) {
	user := &models.User{}
	err := s.db.QueryRowContext(ctx, "SELECT * FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.IsAdmin, &user.ProfilePicture)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *SQLiteUserStore) UpdateUserProfile(ctx context.Context, user *models.User) (*models.User, error) {
	_, err := s.db.ExecContext(ctx, "UPDATE users SET username = ?, email = ? WHERE id = ?", user.Username, user.Email, user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *SQLiteUserStore) UpdateProfilePicture(ctx context.Context, id int, profilePicture string) error {
	_, err := s.db.ExecContext(ctx, "UPDATE users SET profile_picture = ? WHERE id = ?", profilePicture, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SQLiteUserStore) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	rows, err := s.db.QueryContext(ctx, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		var profilePicture sql.NullString
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.IsAdmin, &profilePicture)
		if err != nil {
			return nil, err
		}
		if profilePicture.Valid {
			user.ProfilePicture = &profilePicture.String
		} else {
			user.ProfilePicture = nil
		}
		users = append(users, &user)
	}

	return users, nil
}

func (s *SQLiteUserStore) UpdateUserPassword(ctx context.Context, user *models.User) error {
	_, err := s.db.ExecContext(ctx, "UPDATE users SET password = ? WHERE id = ?", user.Password, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *SQLiteUserStore) DeleteUserByID(ctx context.Context, userID int) (string, error) {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return "", fmt.Errorf("error starting transaction: %v", err)
	}

	defer tx.Rollback() // Rollback if not committed

	// Delete associated jobs first
	_, err = tx.Exec("DELETE FROM jobs WHERE user_id = ?", userID)

	if err != nil {
		return "", fmt.Errorf("error deleting user's jobs: %v", err)
	}

	// Get user's profile picture before deleting user

	var profilePicture sql.NullString
	err = tx.QueryRowContext(ctx, "SELECT profile_picture FROM users WHERE id = ?", userID).Scan(&profilePicture)
	if err != nil {
		return "", fmt.Errorf("error fetching user's profile picture: %v", err)
	}

	// Delete the user
	result, err := tx.ExecContext(ctx, "DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		return "", fmt.Errorf("error deleting user: %v", err)
	}
	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return "", fmt.Errorf("error getting rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return "", sql.ErrNoRows
	}
	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return "", fmt.Errorf("error committing transaction: %v", err)
	}

	return profilePicture.String, nil

}

func (s *SQLiteUserStore) ChangePassword(ctx context.Context, userID int, currentPassword, newPassword string) error {
	// First fetch and validate current password
	var hashedPassword string

	err := s.db.QueryRowContext(ctx, "SELECT password FROM users WHERE id = ?", userID).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user not found")
		}
		return fmt.Errorf("error fetching user password: %v", err)
	}

	// Verify current password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(currentPassword)); err != nil {
		return fmt.Errorf("current password is incorrect")
	}

	// Only if current password is correct, proceed to update
	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing new password: %v", err)
	}

	result, err := s.db.ExecContext(ctx, "UPDATE users SET password = ? WHERE id = ?", hashedNewPassword, userID)
	if err != nil {
		return fmt.Errorf("error updating password: %v", err)
	}

	// Check if update actually affected a row
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking update result: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no user found with id %d", userID)
	}
	return nil
}
