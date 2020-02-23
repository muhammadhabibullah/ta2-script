package model

import (
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	Email     string  `gorm:"column:email"`
	Name      string  `gorm:"column:name"`
	Birthdate string  `gorm:"column:birthdate"`
	Gender    string  `gorm:"column:gender"`
	Weight    float64 `gorm:"column:weight"`
	Height    float64 `gorm:"column:height"`
}
