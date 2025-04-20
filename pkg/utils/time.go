package utils

import "time"

type Duration struct {
	Duration time.Duration `json:"duration"`
	Min      float64       `json:"min"`
	Hours    float64       `json:"hours"`
	Days     float64       `json:"days"`
	Weeks    float64       `json:"weeks"`
	Months   float64       `json:"months"`
	Years    float64       `json:"years"`
}

func NewDuration(duration time.Duration) Duration {
	hours := duration.Hours()

	var (
		years  float64
		months float64
		weeks  float64
		days   float64
	)
	//if hours > 0 {
	days = hours / 24
	weeks = hours / 168
	months = hours / 730
	years = hours / 8760
	//}

	return Duration{
		Duration: duration,
		Min:      duration.Minutes(),
		Hours:    hours,
		Days:     days,
		Weeks:    weeks,
		Months:   months,
		Years:    years,
	}
}
