package model

import "time"

type WeatherResponse struct {
	Location    string    `json:"location"`
	Temp        float64   `json:"temp"`
	Conditions  string    `json:"conditions"`
	RetrievedAt time.Time `json:"retrievedAt"`
	Cached      bool      `json:"cached"`
}
