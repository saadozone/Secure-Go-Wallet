package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/itsmaheshkariya/gin-gorm-rest/config"
	"github.com/itsmaheshkariya/gin-gorm-rest/models"
	"golang.org/x/crypto/bcrypt"
)

// LoginService performs user login and returns the token.
func LoginService(email, password string) *LoginResult {
	var user models.User
	config.DB.First(&user, "email = ?", email)
	if user.ID == 0 {
		return &LoginResult{
			Error: ErrInvalidCredentials,
		}
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return &LoginResult{
			Error: ErrInvalidCredentials,
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte("SECRET"))
	if err != nil {
		return &LoginResult{
			Error: err,
			Token: "",
		}
	}

	return &LoginResult{
		Error: nil,
		Token: tokenString,
	}
}

// LoginResult represents the result of the login service.
type LoginResult struct {
	Error error
	Token string
}

var (
	// ErrInvalidCredentials is returned when the login credentials are invalid.
	ErrInvalidCredentials = &InvalidCredentialsError{}
)

// InvalidCredentialsError represents an error indicating invalid login credentials.
type InvalidCredentialsError struct{}

func (e *InvalidCredentialsError) Error() string {
	return "Invalid email or password"
}
