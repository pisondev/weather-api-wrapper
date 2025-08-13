package service

import (
	"context"
	"encoding/json"
	"pisondev/weather-api-wrapper/cache"
	"pisondev/weather-api-wrapper/client"
	"pisondev/weather-api-wrapper/model"
	"strings"
	"time"
)

type WeatherServiceImpl struct {
	RedisCache    cache.RedisCache
	WeatherClient client.WeatherClient
}

func NewWeatherService(redisCache cache.RedisCache, weatherClient client.WeatherClient) WeatherService {
	return &WeatherServiceImpl{
		RedisCache:    redisCache,
		WeatherClient: weatherClient,
	}
}

func (service *WeatherServiceImpl) GetCurrentConditions(ctx context.Context, location string) (model.Weather, bool, error) {
	cacheKey := "weather:" + strings.ToLower(location)
	cachedData, err := service.RedisCache.Get(ctx, cacheKey)
	if err == nil {
		var weather model.Weather
		err := json.Unmarshal([]byte(cachedData), &weather)
		if err == nil {
			return weather, true, nil
		}
	}

	fetchedData, err := service.WeatherClient.Fetch(ctx, location)
	if err != nil {
		return model.Weather{}, false, err
	}

	weather := model.Weather{
		Location:    fetchedData.ResolvedAddress,
		Temp:        fetchedData.CurentConditions.Temp,
		Conditions:  fetchedData.CurentConditions.Conditions,
		RetrievedAt: time.Now().UTC().Truncate(time.Second),
	}

	weatherJSON, err := json.Marshal(weather)
	if err != nil {
		return model.Weather{}, false, err
	}
	err = service.RedisCache.Set(ctx, cacheKey, weatherJSON, 1*time.Hour)
	if err != nil {
		return model.Weather{}, false, err
	}

	return weather, false, nil
}
