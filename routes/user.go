package routes

import (
	"github.com/gin-gonic/gin"
	// "github.com/itsmaheshkariya/gin-gorm-rest/controller"
	handlers "github.com/itsmaheshkariya/gin-gorm-rest/handler"
)

func UserRoute(router *gin.Engine) {
	// User routes
	router.GET("/users", handlers.GetUsers)
	router.POST("/users", handlers.CreateUsers)
	router.DELETE("/users/:id", handlers.DeleteUsers)
	router.PUT("/users/:id", handlers.UpdateUsers)

	// Laptop routes
	router.GET("/laptops", handlers.GetLapUsers)
	router.GET("/laptops/:id", handlers.GetLapUserByID)
	router.POST("/laptops", handlers.CreateLapUsers)
	router.DELETE("/laptops/:id", handlers.DeleteLapUsers)
	router.PUT("/laptops/:id", handlers.UpdateLapUsers)
}
