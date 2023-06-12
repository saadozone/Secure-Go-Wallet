package server

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/itsmaheshkariya/gin-gorm-rest/handler"
)

// RegisterRoutes registers the routes for the server.
func RegisterRoutes(r *gin.Engine) {
	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.Login)
	r.GET("/validate", handlers.Validate)
}
