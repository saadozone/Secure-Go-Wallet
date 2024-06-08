package routes

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/itsmaheshkariya/gin-gorm-rest/handler"
)

func UserRoute(router *gin.Engine) {
	eh := handlers.NewEmployeeHandler()

	// User routes

	router.GET("/users", eh.GetUsers)
	router.POST("/users", eh.CreateUser)
	router.DELETE("/users/:id", eh.DeleteUser)
	router.PUT("/users/:id", eh.UpdateUser)

	// Employee routes
	router.GET("/employees", eh.GetEmployees)
	router.GET("/employees/:id", eh.GetEmployeeByID)
	router.POST("/employees", eh.CreateEmployee)
	router.DELETE("/employees/:id", eh.DeleteEmployee)
	router.PUT("/employees/:id", eh.UpdateEmployee)
}
