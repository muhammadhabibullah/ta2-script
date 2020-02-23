package model

// Bicycle model
type Bicycle struct {
	ID       int    `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	BikeType string `gorm:"column:biketype"`
	UserID   uint   `gorm:"column:userid"`
}
