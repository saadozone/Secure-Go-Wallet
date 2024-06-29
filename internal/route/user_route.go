package route

import (
	"github.com/gin-gonic/gin"
	"github.com/saadozone/gin-gorm-rest/internal/handler"
	"github.com/saadozone/gin-gorm-rest/internal/middleware"
)

func (r *Router) User(route *gin.RouterGroup, h *handler.Handler) {
	route.GET("/balance", middleware.AuthMiddleware(r.jwtService, r.userService), h.Profile)
}
