package services

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/itsmaheshkariya/gin-gorm-rest/config"
	"github.com/itsmaheshkariya/gin-gorm-rest/models"
)

// ValidateService validates the token and returns the user information.
func ValidateService(c *gin.Context) (*models.User, string, error) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		return nil, "", err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("SECRET"), nil
	})
	if err != nil {
		return nil, "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["sub"].(float64)
		var user models.User
		config.DB.First(&user, int(userID))
		return &user, tokenString, nil
	}

	return nil, "", errors.New("Invalid token")
}
