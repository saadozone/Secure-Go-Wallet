package services

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/itsmaheshkariya/gin-gorm-rest/config"
	"github.com/itsmaheshkariya/gin-gorm-rest/models"
)

// ValidateService validates the token and returns the user information.
func ValidateService(c *gin.Context) (*models.User, string, error) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		return nil, "", errors.New("Token not found in cookie")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetSecretKey()), nil
	})
	if err != nil {
		return nil, "", errors.New("Error parsing token: " + err.Error())
	}

	if !token.Valid {
		return nil, "", errors.New("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, "", errors.New("Invalid token claims")
	}

	userIDFloat, ok := claims["sub"].(float64)
	if !ok {
		return nil, "", errors.New("Invalid user ID in token claims")
	}

	userID := int(userIDFloat)
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return nil, "", fmt.Errorf("Error retrieving user: %w", err)
	}

	return &user, tokenString, nil
}
