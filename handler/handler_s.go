package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/itsmaheshkariya/gin-gorm-rest/models"
	services "github.com/itsmaheshkariya/gin-gorm-rest/service"
)

func GetUsers(c *gin.Context) {
	users := services.GetUsers()
	c.JSON(200, &users)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := services.GetUserByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, &user)
}

func CreateUsers(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	createdUser, err := services.CreateUser(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(200, &createdUser)
}

func DeleteUsers(c *gin.Context) {
	id := c.Param("id")
	err := services.DeleteUser(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, gin.H{"message": "User deleted"})
}

func UpdateUsers(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}
	updatedUser, err := services.UpdateUser(id, &user)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, &updatedUser)
}

func GetLapUsers(c *gin.Context) {
	users := services.GetLapUsers()
	c.JSON(200, &users)
}

func GetLapUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := services.GetLapUserByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, &user)
}
// CreateLapUsers godoc
// @Summary      Create a Laptop User
// @Description  create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        users body models.Laptop true "Laptop"
// @Router       /users [post]
func CreateLapUsers(c *gin.Context) {
	var user models.Laptop
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}
	createdUser, err := services.CreateLapUser(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(200, &createdUser)
}

func DeleteLapUsers(c *gin.Context) {
	id := c.Param("id")
	err := services.DeleteLapUser(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, gin.H{"message": "User deleted"})
}

func UpdateLapUsers(c *gin.Context) {
	var user models.Laptop
	id := c.Param("id")
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}
	updatedUser, err := services.UpdateLapUser(id, &user)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, &updatedUser)
}
