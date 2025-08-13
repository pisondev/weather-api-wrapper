package service

import (
	"context"
	"pisondev/weather-api-wrapper/model"
)

type WeatherService interface {
	GetCurrentConditions(ctx context.Context, location string) (model.Weather, bool, error)
}
