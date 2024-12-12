package services

import (
	"backend/internals/models"
	"backend/internals/store"
	"backend/pkg/utils"
	"context"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	store store.UserStore
}

func NewAuthService(db *sql.DB) *AuthService {
	return &AuthService{store: store.NewUserStore(db)}
}

func (s *AuthService) RegisterUser(ctx context.Context, user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return s.store.CreateUser(ctx, user)
}

func (s *AuthService) LoginUser(ctx context.Context, username, password string) (string, error) {
	user, err := s.store.GetUserByUserName(ctx, username)

	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	return utils.GenerateToken(user.Username, user.ID, user.IsAdmin)
}

func (s *AuthService) ForgotPassword(ctx context.Context, username string) (string, error) {
	user, err := s.store.GetUserByUserName(ctx, username)
	if err != nil {
		return "", err
	}

	generatedPassword := utils.GenerateFromPassword(6)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(generatedPassword), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	user.Password = string(hashedPassword)

	if err := s.store.UpdateUserPassword(ctx, user); err != nil {
		return "", err
	}
	return generatedPassword, nil
}
