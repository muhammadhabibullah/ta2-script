package model

// Raw data model
type Raw struct {
	BicycleID      int     `json:"b"`
	ElapsedSeconds int     `json:"s"`
	Lat            float64 `json:"t"`
	Long           float64 `json:"g"`
	Altitude       float64 `json:"a"`
	Pace           float64 `json:"p"`
	HeartRate      int     `json:"h"`
}
