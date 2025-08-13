package main

import (
	"net/http"
	"os"
	"pisondev/weather-api-wrapper/cache"
	"pisondev/weather-api-wrapper/client"
	"pisondev/weather-api-wrapper/controller"
	"pisondev/weather-api-wrapper/service"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	apiKey := os.Getenv("API_KEY")
	addr := os.Getenv("REDIS_ADDRESS")
	password := os.Getenv("REDIS_PASSWORD")
	dbStr := os.Getenv("REDIS_DB")
	db, err := strconv.Atoi(dbStr)
	if err != nil {
		panic(err)
	}
	httpClient := http.Client{Timeout: 10 * time.Second}

	redisCache := cache.NewRedisCache(addr, password, db)
	weatherClient := client.NewWeatherClient(apiKey, &httpClient)
	weatherService := service.NewWeatherService(redisCache, weatherClient)
	weatherController := controller.NewWeatherController(weatherService)

	app := fiber.New()

	app.Get("/weather/:city", weatherController.GetCurrentConditions)

	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
