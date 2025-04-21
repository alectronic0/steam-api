package utils

import "time"

type Duration struct {
	NanoSec time.Duration `json:"duration"`
	Min     float64       `json:"-"` //`json:"min"`
	Hours   float64       `json:"-"` //`json:"hours"`
	Days    float64       `json:"-"` //`json:"days"`
	Weeks   float64       `json:"-"` //`json:"weeks"`
	Months  float64       `json:"-"` //`json:"months"`
	Years   float64       `json:"-"` //`json:"years"`
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
		NanoSec: duration,
		Min:     duration.Minutes(),
		Hours:   hours,
		Days:    days,
		Weeks:   weeks,
		Months:  months,
		Years:   years,
	}
}
