package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mariombn-io/go-core-api/api/requests"
)

// Auth trata a rota POST /api/auth
func Auth(c *gin.Context) {
	var request requests.AuthRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := request.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.Username == "root" && request.Password == "secret" {
		c.JSON(http.StatusOK, gin.H{"token": "fake-jwt-token"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}
}

func Me(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true})
}
