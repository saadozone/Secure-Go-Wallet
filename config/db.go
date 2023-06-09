package config

import (
	"github.com/itsmaheshkariya/gin-gorm-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres"), &gorm.Config{})
	if err!=nil{
		panic(err)
	}
	err=db.AutoMigrate(&models.User{},&models.Laptop{})
	if err!=nil{
		panic(err)
	}
	DB=db
}