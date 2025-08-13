package client

import "context"

type WeatherClient interface {
	Fetch(ctx context.Context, city string) (*VisualCrossingWeather, error)
}
