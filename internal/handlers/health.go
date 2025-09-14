package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health godoc
// @Summary Health check
// @Description Check if the service is running
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"message": "Service is running",
	})
}