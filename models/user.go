package models

import "gorm.io/gorm"

type Laptop struct {
	gorm.Model
	Year	string		 `json:"year"`
	BrandName string 	 `json:"brandname"`
	ModelName string 	 `json:"model"`
	Price string		 `json:"price"`
}

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string
	Password string
}