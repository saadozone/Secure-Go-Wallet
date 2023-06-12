package stores

import (
	"github.com/itsmaheshkariya/gin-gorm-rest/config"
	"github.com/itsmaheshkariya/gin-gorm-rest/models"
)

// CreateNewUser creates a new user record in the database.
func CreateNewUser(user *models.LapUser) *config.DBResult {
	result := config.DB.Create(user)

	return &config.DBResult{
		Result: result,
		Error:  result.Error,
	}
}

