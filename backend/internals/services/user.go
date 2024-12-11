package services

import (
	"backend/internals/models"
	"backend/internals/store"
	"backend/pkg/utils"
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
)

type UserService struct {
	store store.UserStore
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{store: store.NewUserStore(db)}
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	return s.store.GetUserByID(ctx, id)
}

func (s *UserService) UpdateUserProfile(ctx context.Context, id int, username, emailId string) (*models.User, error) {
	user := &models.User{ID: id, Username: username, Email: emailId}

	return s.store.UpdateUserProfile(ctx, user)
}

func (s *UserService) UpdateProfilePicture(ctx context.Context, id int, profilePicture string) error {
	return s.store.UpdateProfilePicture(ctx, id, profilePicture)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	return s.store.GetAllUsers(ctx)
}

func (s *UserService) DeleteUser(ctx context.Context, userID int) error {
	// Delete user and associated data
	profilePicture, err := s.store.DeleteUserByID(ctx, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user not found")
		}
		return fmt.Errorf("error deleting user: %v", err)
	}

	// Delete profile picture after successful transaction if it exists
	if profilePicture != "" {
		filePath := filepath.Join(os.Getenv("UPLOAD_DIR"), profilePicture)
		err = utils.DeleteFileIfExists(filePath)

		if err != nil {
			return fmt.Errorf("error deleting profile picture: %v", err)
		}
	}

	return nil
}

func (s *UserService) ChangePassword(ctx context.Context, userID int, currentPassword, newPassword string) error {
	return s.store.ChangePassword(ctx, userID, currentPassword, newPassword)
}
