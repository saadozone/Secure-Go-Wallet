package stores

import (
	"github.com/itsmaheshkariya/gin-gorm-rest/config"
	"github.com/itsmaheshkariya/gin-gorm-rest/models"
)

// UserRepository represents the interface for user operations
type UserRepository interface {
	GetUsers() []models.User
	GetUserByID(id string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	DeleteUser(id string) error
	UpdateUser(id string, user *models.User) (*models.User, error)
}

// LaptopRepository represents the interface for laptop operations
type LaptopRepository interface {
	GetLapUsers() []models.Laptop
	GetLapUserByID(id string) (*models.Laptop, error)
	CreateLapUser(user *models.Laptop) (*models.Laptop, error)
	DeleteLapUser(id string) error
	UpdateLapUser(id string, user *models.Laptop) (*models.Laptop, error)
}

type userRepository struct{}

func (r *userRepository) GetUsers() []models.User {
	users := []models.User{}
	config.DB.Find(&users)
	return users
}

func (r *userRepository) GetUserByID(id string) (*models.User, error) {
	user := models.User{}
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) CreateUser(user *models.User) (*models.User, error) {
	result := config.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *userRepository) DeleteUser(id string) error {
	user := models.User{}
	result := config.DB.Delete(&user, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userRepository) UpdateUser(id string, user *models.User) (*models.User, error) {
	existingUser, err := r.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	existingUser.Email = user.Email
	existingUser.Password = user.Password

	result := config.DB.Save(existingUser)
	if result.Error != nil {
		return nil, result.Error
	}
	return existingUser, nil
}

type laptopRepository struct{}

func (r *laptopRepository) GetLapUsers() []models.Laptop {
	users := []models.Laptop{}
	config.DB.Find(&users)
	return users
}

func (r *laptopRepository) GetLapUserByID(id string) (*models.Laptop, error) {
	user := models.Laptop{}
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *laptopRepository) CreateLapUser(user *models.Laptop) (*models.Laptop, error) {
	result := config.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *laptopRepository) DeleteLapUser(id string) error {
	user := models.Laptop{}
	result := config.DB.Delete(&user, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *laptopRepository) UpdateLapUser(id string, user *models.Laptop) (*models.Laptop, error) {
	existingUser, err := r.GetLapUserByID(id)
	if err != nil {
		return nil, err
	}

	existingUser.Year = user.Year
	existingUser.BrandName = user.BrandName
	existingUser.ModelName = user.ModelName
	existingUser.Price = user.Price

	result := config.DB.Save(existingUser)
	if result.Error != nil {
		return nil, result.Error
	}
	return existingUser, nil
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository() UserRepository {
	return &userRepository{}
}

// NewLaptopRepository creates a new LaptopRepository instance
func NewLaptopRepository() LaptopRepository {
	return &laptopRepository{}
}
