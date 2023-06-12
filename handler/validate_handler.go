package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	services "github.com/itsmaheshkariya/gin-gorm-rest/service"
)

// Validate handles the HTTP request for token validation.
func Validate(c *gin.Context) {
	user, tokenString, err := services.ValidateService(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      user,
		"tokenString":  tokenString,
	})
}
