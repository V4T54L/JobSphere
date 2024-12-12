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

	jobHandler := handlers.NewJobHandler(db)
	r.POST("/jobs", jobHandler.CreateJobHandler(db))
	r.GET("/jobsByUser", jobHandler.GetAllJobsByUserHandler(db))
	r.GET("/jobs/:id", jobHandler.GetJobByIdHandler(db))
	r.PUT("/jobs/:id", jobHandler.UpdateJobByHandler(db))
	r.DELETE("/jobs/:id", jobHandler.DeleteJobByHandler(db))

	r.GET("/users", userHandler.GetAllUsersHandler(db))
	r.DELETE("/users/:id", userHandler.DeleteUserByIdHandler(db))
}
