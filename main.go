package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsmaheshkariya/gin-gorm-rest/config"
	handler "github.com/itsmaheshkariya/gin-gorm-rest/handler"
	"github.com/itsmaheshkariya/gin-gorm-rest/middleware"
	"github.com/itsmaheshkariya/gin-gorm-rest/routes"
)
// @title Employee API
// @version 1.0
// @description This is a sample server for managing employees.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
	config.Connect()

	router := gin.Default()

	routes.UserRoute(router)

	router.POST("/signup", handler.Signup)
	router.POST("/login", handler.Login)
	router.GET("/validate", middleware.RequiredAuth, handler.Validate)

	router.Run(":8080")
}
