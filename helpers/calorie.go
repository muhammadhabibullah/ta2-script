package helper

import model "ta2-script/models"

// CountCalorieBurned function
func CountCalorieBurned(cycling model.Finale, user model.User, cyclingTime float64) float64 {
	met := countMET(cycling)
	return met * 3.5 * user.Weight * (cyclingTime / 60) / 200
}

// CountMET function
func countMET(cycling model.Finale) float64 {
	var met float64
	if (cycling.Elevation >= 700) && (cycling.Distance >= 60) {
		met = 14
	} else if cycling.Pace < 16.1 {
		met = 4
	} else if cycling.Pace < 19.2 {
		met = 6.8
	} else if cycling.Pace < 22.4 {
		met = 8
	} else if cycling.Pace < 25.6 {
		met = 10
	} else if cycling.Pace < 30.6 {
		met = 12
	} else if cycling.Pace >= 30.6 {
		met = 15.8
	}
	return met
}
