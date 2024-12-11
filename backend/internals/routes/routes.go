package routes

import (
	"backend/internals/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.GET("/health", handlers.HealthCheckHandler)
}
