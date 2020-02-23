package repository

import (
	database "ta2-script/databases"
	model "ta2-script/models"
)

// GetBicycle return bicycle data
func GetBicycle(bicycleID int) (model.Bicycle, error) {
	var bicycle model.Bicycle
	query := database.DB.Where("id = ?", bicycleID).
		First(&bicycle)

	if query.Error != nil {
		return model.Bicycle{}, query.Error
	}

	return bicycle, nil
}
