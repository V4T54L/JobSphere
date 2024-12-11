package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(ctx *gin.Context) {
	result := map[string]string{
		"status": "healthy",
	}
	ctx.JSON(http.StatusOK, result)
}
