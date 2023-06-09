
package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/itsmaheshkariya/gin-gorm-rest/config"
	"github.com/itsmaheshkariya/gin-gorm-rest/models"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}

func GetUserByID(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	result := config.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, &user)
}

func CreateUsers(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(200, &user)
}

func DeleteUsers(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	result := config.DB.Delete(&user, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, &user)
}

func UpdateUsers(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	result := config.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)
}

func GetLapUsers(c *gin.Context) {
	users := []models.Laptop{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}

func GetLapUserByID(c *gin.Context) {
	var user models.Laptop
	id := c.Param("id")
	result := config.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, &user)
}

func CreateLapUsers(c *gin.Context) {
	var user models.Laptop
	c.BindJSON(&user)
	fmt.Println(user)
	config.DB.Create(&user)
	c.JSON(200, &user)
}

func DeleteLapUsers(c *gin.Context) {
	var user models.Laptop
	id := c.Param("id")
	result := config.DB.Delete(&user, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, &user)
}

func UpdateLapUsers(c *gin.Context) {
	var user models.Laptop
	id := c.Param("id")
	result := config.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)
}