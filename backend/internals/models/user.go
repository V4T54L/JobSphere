package models

type User struct {
	ID             int     `json:"id"`
	Username       string  `json:"username"`
	Password       string  `json:"password"`
	Email          string  `json:"email"`
	IsAdmin        bool    `json:"is_admin"`
	ProfilePicture *string `json:"profile_picture"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required"`
}

type PasswordValidation struct {
	MinLength  int
	HasUpper   bool
	HasLower   bool
	HasNumber  bool
	HasSpecial bool
}

type ForgotPasswordRequest struct {
	Username string `json:"username"`
}
