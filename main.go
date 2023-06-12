package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsmaheshkariya/gin-gorm-rest/config"
	handler "github.com/itsmaheshkariya/gin-gorm-rest/handler"
	"github.com/itsmaheshkariya/gin-gorm-rest/middleware"
	"github.com/itsmaheshkariya/gin-gorm-rest/routes"
)

func main() {
	config.Connect()

	router := gin.Default()

	routes.UserRoute(router)

	router.POST("/signup", handler.Signup)
	router.POST("/login", handler.Login)
	router.GET("/validate", middleware.RequiredAuth, handler.Validate)

	router.Run(":8080")
}
