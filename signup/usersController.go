package signup

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/itsmaheshkariya/gin-gorm-rest/config"
	"github.com/itsmaheshkariya/gin-gorm-rest/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context){
	//get the email/password of the ref body

	var body struct{
		Email string
		Password string
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	
	//hash the password

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password),10)

	if err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	
	//send it back
	user := models.User{Email: body.Email, Password: string(hash)}
	result := config.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}
	//respond
	c.JSON(http.StatusOK, gin.H{})
}
func Login(c *gin.Context){

	//get the email/password of the ref body

	var body struct{
		Email string
		Password string
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	
	//look up requested user
	var user models.User
	config.DB.First(&user, "email=?", body.Email)

	if user.ID == 0{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Invalid user or password",
		})
		return
	}

	//compare sent in password with saved user password hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or Password",
		})
		return
	}

	// Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "sub": user.ID,
    "exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte("SECRET"))

	if err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}
	//SEND IT BACK
	// c.SetSameSite(http.SameSiteLaxMode)
	// c.SetCookie("Authorization", tokenString, 3600*60*30, "", "", false, true)
	
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	tokenString, _ := c.Cookie("Authorization") 
	c.JSON(http.StatusOK, gin.H{
		"message":      user,
		"tokenString":  tokenString,
	})
}

	
