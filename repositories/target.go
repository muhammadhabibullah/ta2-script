package repository

import (
	database "ta2-script/databases"
	model "ta2-script/models"
)

// GetLastestTarget return lastest target
func GetLastestTarget(userID uint) (model.Target, error) {
	var target model.Target
	query := database.DB.Where("userid = ?", userID).
		Last(&target)
	if query.Error != nil {
		return model.Target{}, query.Error
	}

	return target, nil
}
