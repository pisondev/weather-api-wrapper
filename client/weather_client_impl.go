package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type WeatherClientImpl struct {
	ApiKey     string
	HttpClient *http.Client
}

func (client *WeatherClientImpl) Fetch(ctx context.Context, city string) (*VisualCrossingWeather, error) {
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?unitGroup=metric&key=%s&contentType=json&include=current", city, client.ApiKey)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("external api returned an error:%s", resp.Status)
	}

	var visualCrossingWeather VisualCrossingWeather
	err = json.NewDecoder(resp.Body).Decode(&visualCrossingWeather)
	if err != nil {
		return nil, err
	}
	return &visualCrossingWeather, nil
}
