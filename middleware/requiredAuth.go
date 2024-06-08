package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/itsmaheshkariya/gin-gorm-rest/config"
	"github.com/itsmaheshkariya/gin-gorm-rest/models"
)

func RequiredAuth( c *gin.Context){

	// Get the cookie off request
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// Decode it 
		//parse takes the token string and a function for looking up the key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return ([]byte("SECRET")), nil
		
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if token==nil{
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	fmt.Printf("%#v\n",token)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
		// check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)

		}
	
	    // Find the user with token
		var user models.User
		config.DB.First(&user, claims["sub"])

		if user.ID == 0{
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	   // attach to req

	   c.Set("user", user)

	   //Continue

	   c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	
}