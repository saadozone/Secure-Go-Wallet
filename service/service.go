package services

import (
	"github.com/itsmaheshkariya/gin-gorm-rest/models"
	stores "github.com/itsmaheshkariya/gin-gorm-rest/store"
)

func GetUsers() []models.User {
	userRepo := stores.NewUserRepository()
	return userRepo.GetUsers()
}

func GetUserByID(id string) (*models.User, error) {
	userRepo := stores.NewUserRepository()
	return userRepo.GetUserByID(id)
}

func CreateUser(user *models.User) (*models.User, error) {
	userRepo := stores.NewUserRepository()
	return userRepo.CreateUser(user)
}

func DeleteUser(id string) error {
	userRepo := stores.NewUserRepository()
	return userRepo.DeleteUser(id)
}

func UpdateUser(id string, user *models.User) (*models.User, error) {
	userRepo := stores.NewUserRepository()
	return userRepo.UpdateUser(id, user)
}

func GetLapUsers() []models.Laptop {
	laptopRepo := stores.NewLaptopRepository()
	return laptopRepo.GetLapUsers()
}

func GetLapUserByID(id string) (*models.Laptop, error) {
	laptopRepo := stores.NewLaptopRepository()
	return laptopRepo.GetLapUserByID(id)
}

func CreateLapUser(user *models.Laptop) (*models.Laptop, error) {
	laptopRepo := stores.NewLaptopRepository()
	return laptopRepo.CreateLapUser(user)
}

func DeleteLapUser(id string) error {
	laptopRepo := stores.NewLaptopRepository()
	return laptopRepo.DeleteLapUser(id)
}

func UpdateLapUser(id string, user *models.Laptop) (*models.Laptop, error) {
	laptopRepo := stores.NewLaptopRepository()
	return laptopRepo.UpdateLapUser(id, user)
}
