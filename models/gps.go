package model

// GPSRawData model
type GPSRawData struct {
	ID        int     `gorm:"column:id"`
	CyclingID int     `gorm:"column:cyclingid"`
	Seconds   int     `gorm:"column:s"`
	Lat       float64 `gorm:"column:lt"`
	Long      float64 `gorm:"column:lg"`
}