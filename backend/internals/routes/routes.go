package routes

import (
	"backend/internals/handlers"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, db *sql.DB) {
	r.GET("/health", handlers.HealthCheckHandler)

	userHandler := handlers.NewUserHandler(db)
	r.GET("/users/:id", userHandler.GetUserByIdHandler(db))
	r.PUT("/users/:id", userHandler.UpdateUserProfileHandler(db))
	r.POST("/users/:id/picture", userHandler.UpdateUserProfilePcitureHandler(db))
	r.PUT("users/change-password", userHandler.ChangePasswordHandler(db))
}
