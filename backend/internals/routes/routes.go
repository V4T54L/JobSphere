package routes

import (
	"backend/internals/handlers"
	"backend/internals/middlewares"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, db *sql.DB) {
	r.GET("/health", handlers.HealthCheckHandler)

	authHandler := handlers.NewAuthHandler(db)
	jobHandler := handlers.NewJobHandler(db)
	userHandler := handlers.NewUserHandler(db)

	r.POST("/login", authHandler.LoginHandler(db))
	r.POST("/register", authHandler.RegisterHandler(db))
	r.GET("/jobs", jobHandler.GetAllJobsHandler(db))
	r.POST("forgotpassword", authHandler.ForgotPasswordHandler(db))

	authenticatedRoutes := r.Group("/")
	authenticatedRoutes.Use(middlewares.AuthMiddleware())
	authenticatedRoutes.GET("/users/:id", userHandler.GetUserByIdHandler(db))
	authenticatedRoutes.PUT("/users/:id", userHandler.UpdateUserProfileHandler(db))
	authenticatedRoutes.POST("/users/:id/picture", userHandler.UpdateUserProfilePcitureHandler(db))
	authenticatedRoutes.PUT("users/change-password", middlewares.PasswordValidationMiddleware(), userHandler.ChangePasswordHandler(db))

	authenticatedRoutes.POST("/jobs", jobHandler.CreateJobHandler(db))
	authenticatedRoutes.GET("/jobsByUser", jobHandler.GetAllJobsByUserHandler(db))
	authenticatedRoutes.GET("/jobs/:id", jobHandler.GetJobByIdHandler(db))
	authenticatedRoutes.PUT("/jobs/:id", jobHandler.UpdateJobByHandler(db))
	authenticatedRoutes.DELETE("/jobs/:id", jobHandler.DeleteJobByHandler(db))

	authenticatedRoutes.GET("/users", userHandler.GetAllUsersHandler(db))
	authenticatedRoutes.DELETE("/users/:id", userHandler.DeleteUserByIdHandler(db))
}
