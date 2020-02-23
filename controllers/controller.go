package controller

import (
	"encoding/json"
	"fmt"

	model "ta2-script/models"
	service "ta2-script/services"

	"github.com/fatih/color"
)

// CreateRawData controller
func CreateRawData(payload []byte) error {
	var cyclingRawData model.Raw
	if err := json.Unmarshal(payload, &cyclingRawData); err != nil {
		redOutput := color.New(color.FgRed)
		errorOutput := redOutput.Add(color.Bold)
		errorOutput.Println(fmt.Errorf("%s", err))
		return nil
	}

	if err := service.CreateRawData(cyclingRawData); err != nil {
		redOutput := color.New(color.FgRed)
		errorOutput := redOutput.Add(color.Bold)
		errorOutput.Println(fmt.Errorf("%s", err))
		return nil
	}

	return nil
}

// CreateFinaleData controller
func CreateFinaleData(payload []byte) error {
	var cyclingFinaleData model.Finale
	if err := json.Unmarshal(payload, &cyclingFinaleData); err != nil {
		redOutput := color.New(color.FgRed)
		errorOutput := redOutput.Add(color.Bold)
		errorOutput.Println(fmt.Errorf("%s", err))
		return nil
	}

	if err := service.CreateFinaleData(cyclingFinaleData); err != nil {
		redOutput := color.New(color.FgRed)
		errorOutput := redOutput.Add(color.Bold)
		errorOutput.Println(fmt.Errorf("%s", err))
		return nil
	}

	return nil
}
