package handler

import s "github.com/saadozone/gin-gorm-rest/internal/service"

type Handler struct {
	userService        s.UserService
	authService        s.AuthService
	walletService      s.WalletService
	transactionService s.TransactionService
	jwtService         s.JWTService
}

type HandlerConfig struct {
	UserService        s.UserService
	AuthService        s.AuthService
	WalletService      s.WalletService
	TransactionService s.TransactionService
	JWTService         s.JWTService
}

func NewHandler(c *HandlerConfig) *Handler {
	return &Handler{
		userService:        c.UserService,
		authService:        c.AuthService,
		walletService:      c.WalletService,
		transactionService: c.TransactionService,
		jwtService:         c.JWTService,
	}
}
