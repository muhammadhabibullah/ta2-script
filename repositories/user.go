package repository

import (
	database "ta2-script/databases"
	model "ta2-script/models"
)

// GetUser return user data
func GetUser(userID uint) (model.User, error) {
	var user model.User
	query := database.DB.Where("id = ?", userID).
		First(&user)

	if query.Error != nil {
		return model.User{}, query.Error
	}

	return user, nil
}
