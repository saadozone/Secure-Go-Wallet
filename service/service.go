package services

import (
	"github.com/itsmaheshkariya/gin-gorm-rest/models"
	stores "github.com/itsmaheshkariya/gin-gorm-rest/store"
)

type EmployeeService struct {
	EmplpoyeeRepository stores.EmployeeRepository
	UserRepository      stores.UserRepository
}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{
		EmplpoyeeRepository: stores.NewEmployeeRepository(),
		UserRepository:      stores.NewUserRepository(),
	}
}

type Service interface {
	GetEmployees(page, pageSize int) ([]models.Employee, error)
	GetEmployeeByID(id string) (*models.Employee, error)
	CreateEmployee(employee *models.Employee) (*models.Employee, error)
	DeleteEmployee(id string) error
	UpdateEmployee(id string, employee *models.Employee) (*models.Employee, error)
	GetUsers() []models.User
	GetUserByID(id string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	DeleteUser(id string) error
	UpdateUser(id string, user *models.User) (*models.User, error)
}

func (s *EmployeeService) GetUsers() []models.User {
	return s.UserRepository.GetUsers()
}

func (s *EmployeeService) GetUserByID(id string) (*models.User, error) {
	userRepo := stores.NewUserRepository()
	return userRepo.GetUserByID(id)
}

func (s *EmployeeService) CreateUser(user *models.User) (*models.User, error) {
	return s.UserRepository.CreateUser(user)
}

func (s *EmployeeService) DeleteUser(id string) error {
	return s.UserRepository.DeleteUser(id)
}

func (s *EmployeeService) UpdateUser(id string, user *models.User) (*models.User, error) {
	return s.UserRepository.UpdateUser(id, user)
}

func (s *EmployeeService) GetEmployees(page, pageSize int) ([]models.Employee, error) {
	return s.EmplpoyeeRepository.GetEmployees(page, pageSize)
}

func (s *EmployeeService) GetEmployeeByID(id string) (*models.Employee, error) {
	return s.EmplpoyeeRepository.GetEmployeeByID(id)
}

func (s *EmployeeService) CreateEmployee(employee *models.Employee) (*models.Employee, error) {
	return s.EmplpoyeeRepository.CreateEmployee(employee)
}

func (s *EmployeeService) DeleteEmployee(id string) error {
	return s.EmplpoyeeRepository.DeleteEmployee(id)
}

func (s *EmployeeService) UpdateEmployee(id string, employee *models.Employee) (*models.Employee, error) {
	return s.EmplpoyeeRepository.UpdateEmployee(id, employee)
}
