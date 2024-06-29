package route

import (
	"github.com/gin-gonic/gin"
	"github.com/saadozone/gin-gorm-rest/internal/handler"
	"github.com/saadozone/gin-gorm-rest/internal/middleware"
)

func (r *Router) Transaction(route *gin.RouterGroup, h *handler.Handler) {
	route.Use(middleware.AuthMiddleware(r.jwtService, r.userService))
	route.GET("/transactions", h.GetTransactions)
	route.POST("/topup", h.TopUp)
	route.POST("/deduct", h.Transfer)
}
