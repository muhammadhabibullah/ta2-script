package helper

import (
	"time"

	model "ta2-script/models"
)

// CountCyclingDuration return cycling duration in seconds
func CountCyclingDuration(cycling model.Finale) float64 {
	cyclingDuration, _ := time.ParseDuration(cycling.Finishtime.Sub(cycling.Starttime).String())
	return cyclingDuration.Seconds()
}
