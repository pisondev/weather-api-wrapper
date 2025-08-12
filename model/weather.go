package model

import "time"

type Weather struct {
	Location    string
	Temp        float64
	Conditions  string
	RetrievedAt time.Time
}
