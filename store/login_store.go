package stores

import (
	"github.com/itsmaheshkariya/gin-gorm-rest/config"
	"github.com/itsmaheshkariya/gin-gorm-rest/models"
)

// FindUserByEmail finds a user record by email in the database.
func FindUserByEmail(email string) *models.LapUser {
	var user models.LapUser
	config.DB.First(&user, "email = ?", email)
	return &user
}
