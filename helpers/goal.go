package helper

import model "ta2-script/models"

// CountPercentOfGoal count percent of goal
func CountPercentOfGoal(cycling model.Finale, target model.Target, cyclingTime float64, calorieBurned float64) float64 {
	var pog float64
	switch target.TargetType {
	case "D":
		pog = cycling.Distance * 100 / float64(target.TargetNumber)
	case "E":
		pog = cycling.Elevation * 100 / float64(target.TargetNumber)
	case "T":
		pog = float64(cyclingTime) * 100 / ((float64(target.TargetNumber)) * 60)
	case "P":
		pog = cycling.Pace * 100 / float64(target.TargetNumber)
	case "K":
		pog = calorieBurned * 100 / float64(target.TargetNumber)
	}
	return pog
}
