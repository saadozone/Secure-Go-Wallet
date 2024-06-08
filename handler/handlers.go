package handlers

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/itsmaheshkariya/gin-gorm-rest/models"
	services "github.com/itsmaheshkariya/gin-gorm-rest/service"
)

type EmployeeHandler struct {
	EmployeeService services.Service
}

func NewEmployeeHandler() *EmployeeHandler {
	return &EmployeeHandler{
		EmployeeService: services.NewEmployeeService(),
	}
}

// GetUsers handles GET requests to fetch all users
func (e *EmployeeHandler) GetUsers(c *gin.Context) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		users := e.EmployeeService.GetUsers()
		c.JSON(http.StatusOK, users)
	}()

	wg.Wait()
}

// GetUserByID handles GET requests to fetch a user by ID
func (e *EmployeeHandler) GetUserByID(c *gin.Context) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		id := c.Param("id")
		user, err := e.EmployeeService.GetUserByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, user)
	}()

	wg.Wait()
}

// CreateUser handles POST requests to create a new user
func (e *EmployeeHandler) CreateUser(c *gin.Context) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newUser, err := e.EmployeeService.CreateUser(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, newUser)
	}()

	wg.Wait()
}

// DeleteUser handles DELETE requests to delete a user by ID
func (e *EmployeeHandler) DeleteUser(c *gin.Context) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		id := c.Param("id")
		if err := e.EmployeeService.DeleteUser(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	}()

	wg.Wait()
}

// UpdateUser handles PUT requests to update a user by ID
func (e *EmployeeHandler) UpdateUser(c *gin.Context) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		id := c.Param("id")
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		updatedUser, err := e.EmployeeService.UpdateUser(id, &user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, updatedUser)
	}()

	wg.Wait()
}

// GetEmployees handles GET requests to fetch all employees
func (e *EmployeeHandler) GetEmployees(c *gin.Context) {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	pageSizeStr := c.Query("pageSize")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		employees, err := e.EmployeeService.GetEmployees(page, pageSize)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, employees)
	}()

	wg.Wait()
}

// GetEmployeeByID handles GET requests to fetch an employee by ID
func (e *EmployeeHandler) GetEmployeeByID(c *gin.Context) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		id := c.Param("id")
		employee, err := e.EmployeeService.GetEmployeeByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
			return
		}
		c.JSON(http.StatusOK, employee)
	}()

	wg.Wait()
}

// CreateEmployee handles POST requests to create a new employee
func (e *EmployeeHandler) CreateEmployee(c *gin.Context) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		var employee models.Employee
		if err := c.ShouldBindJSON(&employee); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newEmployee, err := e.EmployeeService.CreateEmployee(&employee)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, newEmployee)
	}()

	wg.Wait()
}

// DeleteEmployee handles DELETE requests to delete an employee by ID
func (e *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		id := c.Param("id")
		if err := e.EmployeeService.DeleteEmployee(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	}()

	wg.Wait()
}

// UpdateEmployee handles PUT requests to update an employee by ID

func (e *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		id := c.Param("id")
		var employee models.Employee
		if err := c.ShouldBindJSON(&employee); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		updatedEmployee, err := e.EmployeeService.UpdateEmployee(id, &employee)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, updatedEmployee)
	}()

	wg.Wait()
}
