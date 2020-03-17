package repository

import (
	database "ta2-script/databases"
	model "ta2-script/models"
)

// GetLatestCyclingByBicycleID return cycling data
func GetLatestCyclingByBicycleID(bicycleID int) (model.Cycling, error) {
	var cycling model.Cycling
	query := database.DB.Where("bicycleid = ?", bicycleID).
		Last(&cycling)
	if query.Error != nil {
		return model.Cycling{}, query.Error
	}

	return cycling, nil
}

// InsertCycling save cycling data
func InsertCycling(cycling model.Cycling) error {
	return database.DB.Save(&cycling).Error
}
