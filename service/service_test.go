package services_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mock_store "github.com/itsmaheshkariya/gin-gorm-rest/mocks"
	"github.com/itsmaheshkariya/gin-gorm-rest/models"
	services "github.com/itsmaheshkariya/gin-gorm-rest/service"
)

func TestNewEmployeeService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := services.NewEmployeeService()

	// Assert that the returned service is not nil
	assert.NotNil(t, service, "expected a non-nil service object")

}

func TestEmployeeService_GetEmployeeByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_store.NewMockEmployeeRepository(ctrl)
	service := services.EmployeeService{
		EmplpoyeeRepository: mockRepo,
	}

	mockEmployee := &models.Employee{
		Position: "sde2",
		Name:     "John Doe",
		Salary:   23000,
	}

	// Specify the expectations for the mock repository method call
	mockRepo.EXPECT().GetEmployeeByID("1").Return(mockEmployee, nil)

	// Call the function under test
	employee, err := service.GetEmployeeByID("1")

	// Check the returned employee and error
	assert.NotNil(t, employee, "expected a non-nil employee object")
	assert.NoError(t, err, "expected no error")
}
func TestEmployeeService_GetEmployees(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_store.NewMockEmployeeRepository(ctrl)
	service := services.EmployeeService{
		EmplpoyeeRepository: mockRepo,
	}

	// Prepare a mock response
	mockEmployees := []models.Employee{
		{Position: "sde1", Name: "John Doe"},
		{Position: "sde2", Name: "Jane 3 Doe"},
		// Add more mock employees as needed
	}

	// Specify the expectations for the mock repository method call
	page := 1
	pageSize := 10
	mockRepo.EXPECT().GetEmployees(page, pageSize).Return(mockEmployees, nil)

	// Call the function under test
	employees, err := service.GetEmployees(page, pageSize)

	// Check the returned employees and error
	assert.NotNil(t, employees, "expected non-nil employee slice")
	assert.NoError(t, err, "expected no error")
	assert.Equal(t, len(mockEmployees), len(employees), "expected length of returned employees to match mock data")
}
func TestEmployeeService_CreateEmployee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_store.NewMockEmployeeRepository(ctrl)
	service := services.EmployeeService{
		EmplpoyeeRepository: mockRepo,
	}

	// Prepare a mock employee to be created
	mockEmployee := &models.Employee{
		Name:     "John Doe",
		Position: "sde4"}

	// Specify the expectations for the mock repository method call
	mockCreatedEmployee := &models.Employee{
		Position: "sde1",
		Name:     "John Doe",
		Salary:   40000}
	mockRepo.EXPECT().CreateEmployee(mockEmployee).Return(mockCreatedEmployee, nil)

	createdEmployee, err := service.CreateEmployee(mockEmployee)

	assert.NotNil(t, createdEmployee, "expected non-nil created employee")
	assert.NoError(t, err, "expected no error")
	assert.Equal(t, mockCreatedEmployee.ID, createdEmployee.ID, "expected ID of created employee to match mock data")
	assert.Equal(t, mockCreatedEmployee.Name, createdEmployee.Name, "expected name of created employee to match mock data")
}
func TestEmployeeService_DeleteEmployee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_store.NewMockEmployeeRepository(ctrl)
	service := services.EmployeeService{
		EmplpoyeeRepository: mockRepo,
	}

	// Specify the ID of the employee to be deleted
	id := "1"

	// Specify the expectations for the mock repository method call
	mockRepo.EXPECT().DeleteEmployee(id).Return(nil)

	// Call the function under test
	err := service.DeleteEmployee(id)

	// Check the error returned
	assert.NoError(t, err, "expected no error")
}
func TestEmployeeService_UpdateEmployee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_store.NewMockEmployeeRepository(ctrl)
	service := services.EmployeeService{
		EmplpoyeeRepository: mockRepo,
	}

	// Specify the ID of the employee to be updated
	id := "1"

	// Prepare a mock employee with updated details
	updatedEmployee := &models.Employee{
		Position: "sde1",
		Name:     "Updated Name",
		Salary:   50000}

	// Specify the expectations for the mock repository method call
	mockRepo.EXPECT().UpdateEmployee(id, updatedEmployee).Return(updatedEmployee, nil)

	// Call the function under test
	returnedEmployee, err := service.UpdateEmployee(id, updatedEmployee)

	// Check the returned updated employee and error
	assert.NotNil(t, returnedEmployee, "expected non-nil updated employee")
	assert.NoError(t, err, "expected no error")
	assert.Equal(t, updatedEmployee.ID, returnedEmployee.ID, "expected ID of updated employee to match mock data")
	assert.Equal(t, updatedEmployee.Name, returnedEmployee.Name, "expected name of updated employee to match mock data")
}
