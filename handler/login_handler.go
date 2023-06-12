package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	services "github.com/itsmaheshkariya/gin-gorm-rest/service"
)

// Login handles the HTTP request for user login.
func Login(c *gin.Context) {
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

	result := services.LoginService(body.Email, body.Password)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": result.Token,
	})
}
