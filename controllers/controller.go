package controller

import (
	"encoding/json"

	logger "ta2-script/loggers"
	model "ta2-script/models"
	service "ta2-script/services"
)

// CreateRawData controller
func CreateRawData(payload []byte) error {
	var cyclingRawData model.Raw
	if err := json.Unmarshal(payload, &cyclingRawData); err != nil {
		logger.LogRedError(err)
		return nil
	}

	if err := service.CreateRawData(cyclingRawData); err != nil {
		logger.LogRedError(err)
		return nil
	}

	return nil
}

// CreateFinaleData controller
func CreateFinaleData(payload []byte) error {
	var cyclingFinaleData model.Finale
	if err := json.Unmarshal(payload, &cyclingFinaleData); err != nil {
		logger.LogRedError(err)
		return nil
	}

	if err := service.CreateFinaleData(cyclingFinaleData); err != nil {
		logger.LogRedError(err)
		return nil
	}

	return nil
}
