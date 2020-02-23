package model

import (
	"github.com/jinzhu/gorm"
)

// Target model
type Target struct {
	gorm.Model
	Name         string `gorm:"column:name"`
	TargetType   string `gorm:"column:targettype"`
	TargetNumber int    `gorm:"column:targetnumber"`
	UserID       int    `gorm:"column:userid"`
}
