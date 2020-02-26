package helper

import (
	"fmt"
	"math"
	"strings"

	model "ta2-script/models"
)

// GetRecommendation return cycling recommendation
func GetRecommendation(cycling model.Finale, user model.User, target model.Target, goalPercent float64) string {
	var recommendations []string

	userAge, err := CountAge(user.Birthdate)
	if err != nil {
		return ""
	}

	heartRateRecommendation := getHeartRateRecommendation(userAge, cycling.HeartRate)
	recommendations = append(recommendations, heartRateRecommendation)

	if target.TargetType == "K" {
		dailyCalorieIntakeRecommendation := getDailyCalorieIntakeRecommendation(cycling, user, userAge)
		recommendations = append(recommendations, dailyCalorieIntakeRecommendation)
	} else if target.TargetType != "" {
		targetRecommendation := getTargetRecommendation(cycling, target, goalPercent)
		recommendations = append(recommendations, targetRecommendation)
	}

	allRecommendations := strings.Join(recommendations, " ")
	return allRecommendations
}

// getHeartRateRecommendation function
func getHeartRateRecommendation(age int, BPM int) string {
	maxBPM := 220 - age
	if float64(BPM) <= (0.5 * float64(maxBPM)) {
		return "Rata-rata detak jantung Anda di bawah 50%. Tingkatkan lagi olahraga Anda!"
	} else if float64(BPM) <= (0.6 * float64(maxBPM)) {
		return "Rata-rata detak jantung Anda berada pada kisaran 50-60% detak jantung maksimal. Anda sudah meningkatkan kebugaran tubuh Anda."
	} else if float64(BPM) <= (0.7 * float64(maxBPM)) {
		return "Rata-rata detak jantung Anda berada pada kisaran 60-70% detak jantung maksimal. Anda sudah meningkatkan ketahanan tubuh dan pembakaran kalori Anda."
	} else if float64(BPM) <= (0.8 * float64(maxBPM)) {
		return "Rata-rata detak jantung Anda berada pada kisaran 70-80% detak jantung maksimal. Anda sudah meningkatkan kapasitas aerobik Anda."
	} else if float64(BPM) <= (0.9 * float64(maxBPM)) {
		return "Rata-rata detak jantung Anda berada pada kisaran 80-90% detak jantung maksimal. Anda sudah meningkatkan batas performa tubuh Anda."
	} else {
		return "Rata-rata detak jantung Anda berada pada kisaran 90-100% detak jantung maksimal. Anda sudah mencapai performa maksimal tubuh Anda. Jangan paksakan lebih lanjut!"
	}
}

// getDailyCalorieIntakeRecommendation function
func getDailyCalorieIntakeRecommendation(cycling model.Finale, user model.User, age int) string {
	dailyCalorieIntake := countDCI(countBMR(user, age), measureActivityLevel(countMET(cycling)))
	return "Penuhi asupan kalori sebanyak " + fmt.Sprintf("%.2f", math.Round((dailyCalorieIntake*100)/100)) +
		" kcal hari ini untuk menjaga berat badan."
}

// measureActivityLevel function
func measureActivityLevel(met float64) string {
	if met <= 6 {
		return "Moderate"
	}
	return "High"
}

// countBMR function
func countBMR(user model.User, age int) float64 {
	if user.Gender == "M" {
		return 66.47 + (13.75 * user.Weight) + (5.003 * user.Height) - (6.755 * float64(age))
	}
	return 655.1 + (9.563 * user.Weight) + (1.850 * user.Height) - (4.676 * float64(age)) // Female
}

// countDCI function
func countDCI(bmr float64, activityLevel string) float64 {
	if activityLevel == "Moderate" {
		return bmr * 1.55
	}
	return bmr * 1.7 // High activity level
}

func getTargetRecommendation(cycling model.Finale, target model.Target, goalPercent float64) string {
	var targetStr, recommendationStr string
	switch target.TargetType {
	case "T":
		targetStr = "waktu"
	case "D":
		targetStr = "total jarak tempuh"
	case "P":
		targetStr = "kecepatan"
	case "E":
		targetStr = "total elevasi"
	}
	if goalPercent < 50 {
		recommendationStr = "Target " + targetStr + " Anda masih di bawah 50%. Tingkatkan lagi intensitas latihan Anda."
	} else if goalPercent < 75 {
		recommendationStr = "Target " + targetStr + " Anda sudah mencapai 50-75%. Jaga semangat Anda untuk mencapai target."
	} else if goalPercent < 90 {
		recommendationStr = "Target " + targetStr + " Anda sudah mencapai 75-90%. Jangan mudah menyerah."
	} else if goalPercent < 100 {
		recommendationStr = "Target " + targetStr + " Anda sudah di atas 90%. Sedikit lagi dan capai target Anda."
	} else {
		recommendationStr = "Target " + targetStr + " Anda sudah tercapai. Selamat! Perbarui target jika Anda butuh tantangan baru."
	}
	return recommendationStr
}
