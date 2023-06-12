package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	services "github.com/itsmaheshkariya/gin-gorm-rest/service"
)

// Signup handles the HTTP request for user signup.
func Signup(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	result := services.SignupService(body.Email, body.Password)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
			"cause": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
