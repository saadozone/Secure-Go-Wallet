package utils

import (
	"net/http"

	"github.com/saadozone/gin-gorm-rest/pkg/custom_error"
)

func GetStatusCode(err error) int {
	var statusCode int = http.StatusInternalServerError

	if _, ok := err.(*custom_error.NotValidEmailError); ok {
		statusCode = http.StatusUnprocessableEntity
	} else if _, ok := err.(*custom_error.UserAlreadyExistsError); ok {
		statusCode = http.StatusConflict
	} else if _, ok := err.(*custom_error.IncorrectCredentialsError); ok {
		statusCode = http.StatusUnauthorized
	} else if _, ok := err.(*custom_error.UserNotFoundError); ok {
		statusCode = http.StatusBadRequest
	} else if _, ok := err.(*custom_error.PasswordNotSame); ok {
		statusCode = http.StatusUnprocessableEntity
	} else if _, ok := err.(*custom_error.ResetTokenNotFound); ok {
		statusCode = http.StatusBadRequest
	} else if _, ok := err.(*custom_error.SourceOfFundNotFoundError); ok {
		statusCode = http.StatusBadRequest
	} else if _, ok := err.(*custom_error.InsufficientBallanceError); ok {
		statusCode = http.StatusBadRequest
	} else if _, ok := err.(*custom_error.WalletNotFoundError); ok {
		statusCode = http.StatusBadRequest
	} else if _, ok := err.(*custom_error.WalletAlreadyExistsError); ok {
		statusCode = http.StatusConflict
	} else if _, ok := err.(*custom_error.TransferToSameWalletError); ok {
		statusCode = http.StatusBadRequest
	}
	return statusCode
}
