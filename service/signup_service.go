package services

import (
	"github.com/itsmaheshkariya/gin-gorm-rest/config"
	"github.com/itsmaheshkariya/gin-gorm-rest/models"
	"golang.org/x/crypto/bcrypt"
)

// SignupService performs the user signup and returns the result.
func SignupService(email, password string) *config.DBResult {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return &config.DBResult{
			Error: err,
		}
	}

	user := models.User{
		Email:    email,
		Password: string(hash),
	}

	result := config.DB.Create(&user)
	return &config.DBResult{
		Result: result,
		Error:  result.Error,
	}
}
