package helper

import "time"

// CountAge function
func CountAge(birthdate string) (int, error) {
	layout := "2006-01-02T15:04:05Z"
	birthday, err := time.Parse(layout, birthdate)
	if err != nil {
		return 0, err
	}

	now := time.Now()
	years := now.Year() - birthday.Year()
	if now.YearDay() < birthday.YearDay() {
		years--
	}

	return years, nil
}
