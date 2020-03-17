package service

import (
	helper "ta2-script/helpers"
	model "ta2-script/models"
	repository "ta2-script/repositories"

	"github.com/jinzhu/gorm"
)

// CreateRawData service
func CreateRawData(raw model.Raw) error {
	cycling, err := repository.GetLatestCyclingByBicycleID(raw.BicycleID)
	if err != nil {
		return err
	}

	gps := model.GPSRawData{
		CyclingID: cycling.ID,
		Seconds:   raw.ElapsedSeconds,
		Lat:       raw.Lat,
		Long:      raw.Long,
	}
	return repository.InsertCyclingGPS(gps)
}

// CreateFinaleData service
func CreateFinaleData(finale model.Finale) error {
	bicycle, err := repository.GetBicycle(finale.BicycleID)
	if err != nil {
		return err
	}
	user, err := repository.GetUser(bicycle.UserID)
	if err != nil {
		return err
	}
	target, err := repository.GetLastestTarget(user.ID)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err
	}

	cyclingDuration := helper.CountCyclingDuration(finale)
	calorieBurned := helper.CountCalorieBurned(finale, user, cyclingDuration)
	goalPercent := helper.CountPercentOfGoal(finale, target, cyclingDuration, calorieBurned)
	recommendation := helper.GetRecommendation(finale, user, target, goalPercent)

	cycling := model.Cycling{
		Starttime:      finale.Starttime,
		Finishtime:     finale.Finishtime,
		Averagepace:    finale.Pace,
		Elevationgain:  finale.Elevation,
		Distance:       finale.Distance,
		Heartrate:      finale.HeartRate,
		Calorieburned:  calorieBurned,
		Percentofgoal:  goalPercent,
		Recommendation: recommendation,
		Userid:         int(user.ID),
		Bicycleid:      finale.BicycleID,
	}
	return repository.InsertCycling(cycling)
}
