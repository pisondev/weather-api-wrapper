package controller

import (
	"pisondev/weather-api-wrapper/model"
	"pisondev/weather-api-wrapper/service"

	"github.com/gofiber/fiber/v2"
)

type WeatherControllerImpl struct {
	WeatherService service.WeatherService
}

func (controller *WeatherControllerImpl) GetCurrentConditions(ctx *fiber.Ctx) error {
	city := ctx.Params("city")
	weather, isCached, err := controller.WeatherService.GetCurrentConditions(ctx.Context(), city)
	if err != nil {
		return err
	}

	weatherResponse := model.WeatherResponse{
		Location:    weather.Location,
		Temp:        weather.Temp,
		Conditions:  weather.Conditions,
		RetrievedAt: weather.RetrievedAt,
		Cached:      isCached,
	}
	return ctx.Status(fiber.StatusOK).JSON(weatherResponse)
}
