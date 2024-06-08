package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string
	Password string
}
// User represents the user model.
type LapUser struct {
	gorm.Model
	Email    string
	Password string
}
