package handlers

import (
	"backend/internals/models"
	"backend/internals/services"
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const queryTimeout = time.Second * 5

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{service: *services.NewUserService(db)}
}

func (h *UserHandler) GetUserByIdHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		ctx, cancel := context.WithTimeout(c, queryTimeout)
		defer cancel()

		user, err := h.service.GetUserByID(ctx, id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func (h *UserHandler) UpdateUserProfileHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		var userUpdate struct {
			Username string `json:"username"`
			Email    string `json:"email"`
		}

		if err := c.ShouldBindJSON(&userUpdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userID := c.GetInt("userID")
		isAdmin := c.GetBool("isAdmin")

		if !isAdmin && userID != id {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized to update this user profile"})
			return
		}

		ctx, cancel := context.WithTimeout(c, queryTimeout)
		defer cancel()

		updateUser, err := h.service.UpdateUserProfile(ctx, id, userUpdate.Username, userUpdate.Email)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user profile"})
			return
		}

		c.JSON(http.StatusOK, updateUser)
	}
}

func (h *UserHandler) UpdateUserProfilePcitureHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		userID := c.GetInt("userID")
		isAdmin := c.GetBool("isAdmin")

		if !isAdmin && userID != id {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized to update this user profile"})
			return
		}

		file, err := c.FormFile("profile_picture")

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error uploading file"})
			return
		}

		if err := os.MkdirAll(os.Getenv("UPLOAD_DIR"), os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating upload directory"})
			return
		}

		filename := fmt.Sprintf("%d-%s", id, filepath.Base(file.Filename))
		filePath := filepath.Join(os.Getenv("UPLOAD_DIR"), filename)

		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving uploaded file"})
			return
		}

		ctx, cancel := context.WithTimeout(c, queryTimeout)
		defer cancel()

		err = h.service.UpdateProfilePicture(ctx, id, filename)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating profile picture in database"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Profile picture updated successfully"})
	}
}

func (h *UserHandler) GetAllUsersHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin := c.GetBool("isAdmin")
		if !isAdmin {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized to get all users"})
			return
		}

		ctx, cancel := context.WithTimeout(c, queryTimeout)
		defer cancel()

		users, err := h.service.GetAllUsers(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

func (h *UserHandler) DeleteUserByIdHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if user is admin
		isAdmin := c.GetBool("isAdmin")
		if !isAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access"})
			return
		}
		// Get user ID from request params
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		// Check if user is trying to delete themselves
		currentUserID := c.GetInt("userID")
		if currentUserID == id {
			c.JSON(http.StatusBadRequest, gin.H{"error": "You cannot delete yourself"})
			return
		}

		ctx, cancel := context.WithTimeout(c, queryTimeout)
		defer cancel()

		// Delete User
		err = h.service.DeleteUser(ctx, id)
		if err != nil {
			if err.Error() == "user not found" {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error deleting user: %v", err)})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User and associated data deleted successfully"})

	}
}

func (h *UserHandler) ChangePasswordHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.ChangePasswordRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userID := c.GetInt("userID")
		ctx, cancel := context.WithTimeout(c, queryTimeout)
		defer cancel()

		err := h.service.ChangePassword(ctx, userID, req.CurrentPassword, req.NewPassword)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
	}
}
