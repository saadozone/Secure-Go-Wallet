package route

import s "github.com/saadozone/gin-gorm-rest/internal/service"

type Router struct {
	userService s.UserService
	jwtService  s.JWTService
}

type RouterConfig struct {
	UserService s.UserService
	JWTService  s.JWTService
}

func NewRouter(c *RouterConfig) *Router {
	return &Router{
		userService: c.UserService,
		jwtService:  c.JWTService,
	}
}
