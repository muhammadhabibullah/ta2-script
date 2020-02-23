package repository

import (
	database "ta2-script/databases"
	model "ta2-script/models"
)

// InsertCyclingGPS save gps data
func InsertCyclingGPS(gps model.GPSRawData) error {
	return database.DB.Save(&gps).Error
}
