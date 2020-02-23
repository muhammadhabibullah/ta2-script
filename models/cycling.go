package model

import "time"

// Cycling model
type Cycling struct {
	ID             int       `gorm:"column:id"`
	Starttime      time.Time `gorm:"column:starttime"`
	Finishtime     time.Time `gorm:"column:finishtime"`
	Averagepace    float64   `gorm:"column:averagepace"`
	Elevationgain  float64   `gorm:"column:elevationgain"`
	Distance       float64   `gorm:"column:distance"`
	Heartrate      int       `gorm:"column:heartrate"`
	Calorieburned  float64   `gorm:"column:calorieburned"`
	Percentofgoal  float64   `gorm:"column:percentofgoal"`
	Recommendation string    `gorm:"column:recommendation"`
	Userid         int       `gorm:"column:userid"`
	Bicycleid      int       `gorm:"column:bicycleid"`
}
