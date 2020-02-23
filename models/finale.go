package model

import "time"

// Finale data model
type Finale struct {
	BicycleID  int       `json:"b"`
	Starttime  time.Time `json:"s"`
	Finishtime time.Time `json:"f"`
	Distance   float64   `json:"d"`
	Elevation  float64   `json:"e"`
	Pace       float64   `json:"p"`
	HeartRate  int       `json:"h"`
}
