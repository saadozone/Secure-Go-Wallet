package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsmaheshkariya/gin-gorm-rest/config"
	"github.com/itsmaheshkariya/gin-gorm-rest/middleware"
	"github.com/itsmaheshkariya/gin-gorm-rest/routes"
	"github.com/itsmaheshkariya/gin-gorm-rest/signup"
)

func main(){
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.POST("/signup", signup.Signup)
	router.POST("/login", signup.Login)
	router.GET("/validate",middleware.RequiredAuth, signup.Validate)


	router.Run(":8080")
}