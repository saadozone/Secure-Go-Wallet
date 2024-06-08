package handlers_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	handlers "github.com/itsmaheshkariya/gin-gorm-rest/handler"
	"github.com/itsmaheshkariya/gin-gorm-rest/mocks"
	"github.com/itsmaheshkariya/gin-gorm-rest/models"
)

func TestGetEmployeeByIDHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockService(ctrl)
	handler := handlers.EmployeeHandler{
		EmployeeService: mockService,
	}

	// Prepare a mock response
	mockEmployee := &models.Employee{
		Position: "sde1",
		Name:     "John Doe",
	}

	// Specify the expectations for the mock service method call
	mockService.EXPECT().GetEmployeeByID("1").Return(mockEmployee, nil)

	// Create a new gin context for testing
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	// Run the handler function in a goroutine with a wait group
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		handler.GetEmployeeByID(c)
	}()
	wg.Wait()

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "John Doe")
}

func TestGetEmployeeByIDHandler_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockService(ctrl)
	handler := handlers.EmployeeHandler{
		EmployeeService: mockService,
	}

	// Specify the expectations for the mock service method call
	mockService.EXPECT().GetEmployeeByID("1").Return(nil, errors.New("Employee not found"))

	// Create a new gin context for testing
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	// Run the handler function in a goroutine with a wait group
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		handler.GetEmployeeByID(c)
	}()
	wg.Wait()

	// Check the response
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Employee not found")
}
func TestGetEmployeesHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockService(ctrl)
	handler := handlers.EmployeeHandler{
		EmployeeService: mockService,
	}

	// Prepare a mock response
	mockEmployees := []models.Employee{
		{Position: "sde1", Name: "John Doe"},
		{Position: "sde2", Name: "Jane Doe"},
	}

	// Specify the expectations for the mock service method call
	page := 1
	pageSize := 10
	mockService.EXPECT().GetEmployees(page, pageSize).Return(mockEmployees, nil)

	// Create a new gin context for testing
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/employees?page=1&pageSize=10", nil)

	// Run the handler function in a goroutine with a wait group
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		handler.GetEmployees(c)
	}()
	wg.Wait()

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "John Doe")
	assert.Contains(t, w.Body.String(), "Jane Doe")
}

func TestGetEmployeesHandler_InvalidPageAndPageSize(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockService(ctrl)
	handler := handlers.EmployeeHandler{
		EmployeeService: mockService,
	}

	// Prepare a mock response
	mockEmployees := []models.Employee{
		{Position: "sde1", Name: "John Doe"},
		{Position: "sde2", Name: "Jane Doe"},
	}

	// Set default values for page and pageSize
	page := 1
	pageSize := 10
	mockService.EXPECT().GetEmployees(page, pageSize).Return(mockEmployees, nil)

	// Create a new gin context for testing with invalid page and pageSize
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/employees?page=invalid&pageSize=invalid", nil)

	// Run the handler function in a goroutine with a wait group
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		handler.GetEmployees(c)
	}()
	wg.Wait()

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "John Doe")
	assert.Contains(t, w.Body.String(), "Jane Doe")
}

func TestGetEmployeesHandler_InternalServerError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockService(ctrl)
	handler := handlers.EmployeeHandler{
		EmployeeService: mockService,
	}

	// Specify the expectations for the mock service method call
	page := 1
	pageSize := 10
	mockService.EXPECT().GetEmployees(page, pageSize).Return(nil, errors.New("internal server error"))

	// Create a new gin context for testing
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/employees?page=1&pageSize=10", nil)

	// Run the handler function in a goroutine with a wait group
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		handler.GetEmployees(c)
	}()
	wg.Wait()

	// Check the response
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "internal server error")
}
func TestCreateEmployeeHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockService(ctrl)
	handler := handlers.EmployeeHandler{
		EmployeeService: mockService,
	}

	// Prepare a mock request and response
	employee := &models.Employee{
		Name: "John Doe",
	}
	mockEmployee := &models.Employee{
		Position: "1",
		Name:     "John Doe",
	}

	// Specify the expectations for the mock service method call
	mockService.EXPECT().CreateEmployee(employee).Return(mockEmployee, nil)

	// Create a new gin context for testing
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	jsonBody := `{"name":"John Doe"}`
	c.Request = httptest.NewRequest("POST", "/employees", bytes.NewBufferString(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Run the handler function in a goroutine with a wait group
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		handler.CreateEmployee(c)
	}()
	wg.Wait()

	// Check the response
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "John Doe")
}

func TestDeleteEmployeeHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockService(ctrl)
	handler := handlers.EmployeeHandler{
		EmployeeService: mockService,
	}

	// Specify the expectations for the mock service method call
	mockService.EXPECT().DeleteEmployee("1").Return(errors.New("internal server error"))

	// Create a new gin context for testing
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}
	c.Request = httptest.NewRequest("DELETE", "/employees/1", nil)

	// Run the handler function in a goroutine with a wait group
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		handler.DeleteEmployee(c)
	}()
	wg.Wait()

	// Check the response
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "internal server error")
}
func TestUpdateEmployeeHandlert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockService(ctrl)
	handler := handlers.EmployeeHandler{
		EmployeeService: mockService,
	}

	// Create a new gin context for testing
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}
	c.Request = httptest.NewRequest("PUT", "/employees/1", strings.NewReader(`{"name":123,"email":"johndoe@example.com"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	// Run the handler function in a goroutine with a wait group
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		handler.UpdateEmployee(c)
	}()
	wg.Wait()

	// Check the response
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

