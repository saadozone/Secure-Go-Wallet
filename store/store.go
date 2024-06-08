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

// EmployeeRepository represents the interface for employee operations
type EmployeeRepository interface {
	GetEmployees(page, pageSize int) ([]models.Employee, error)
	GetEmployeeByID(id string) (*models.Employee, error)
	CreateEmployee(employee *models.Employee) (*models.Employee, error)
	DeleteEmployee(id string) error
	UpdateEmployee(id string, employee *models.Employee) (*models.Employee, error)
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

type employeeRepository struct{}

func (r *employeeRepository) GetEmployees(page, pageSize int) ([]models.Employee, error) {
	var employees []models.Employee
	offset := (page - 1) * pageSize
	result := config.DB.Offset(offset).Limit(pageSize).Find(&employees)
	if result.Error != nil {
		return nil, result.Error
	}
	return employees, nil
}

func (r *employeeRepository) GetEmployeeByID(id string) (*models.Employee, error) {
	employee := models.Employee{}
	result := config.DB.First(&employee, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &employee, nil
}

func (r *employeeRepository) CreateEmployee(employee *models.Employee) (*models.Employee, error) {
	result := config.DB.Create(employee)
	if result.Error != nil {
		return nil, result.Error
	}
	return employee, nil
}

func (r *employeeRepository) DeleteEmployee(id string) error {
	employee := models.Employee{}
	result := config.DB.Delete(&employee, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *employeeRepository) UpdateEmployee(id string, employee *models.Employee) (*models.Employee, error) {
	existingEmployee, err := r.GetEmployeeByID(id)
	if err != nil {
		return nil, err
	}

	existingEmployee.Name = employee.Name
	existingEmployee.Position = employee.Position
	existingEmployee.Salary = employee.Salary

	result := config.DB.Save(existingEmployee)
	if result.Error != nil {
		return nil, result.Error
	}
	return existingEmployee, nil
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository() UserRepository {
	return &userRepository{}
}

// NewEmployeeRepository creates a new EmployeeRepository instance
func NewEmployeeRepository() EmployeeRepository {
	return &employeeRepository{}
}
