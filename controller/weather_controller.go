package controller

import "github.com/gofiber/fiber/v2"

type WeatherController interface {
	GetCurrentConditions(ctx *fiber.Ctx) error
}
