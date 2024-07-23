package server

import (
	"module/internal/models"
	"module/internal/services"

	"github.com/gofiber/fiber/v2"
)

func carInsert(c *fiber.Ctx) error {

	level, res := services.TokenCheck(c)

	if level < 2 {
		return models.ResponseAccessDenied()
	}

	if res.Error() == models.ResponseTokenGood().Error() {
		return services.SendcarInsert(c)
	}

	return models.ResponseTokenExpired()
}

func carShow(c *fiber.Ctx) error {

	level, res := services.TokenCheck(c)

	if level < 1 {
		return models.ResponseAccessDenied()
	}

	if res.Error() == models.ResponseTokenGood().Error() {
		return services.SendcarShow(c)
	}

	return models.ResponseTokenExpired()
}

func carUpdate(c *fiber.Ctx) error {

	level, res := services.TokenCheck(c)

	if level < 2 {
		return models.ResponseAccessDenied()
	}

	if res.Error() == models.ResponseTokenGood().Error() {
		return services.SendcarUpdate(c)
	}

	return models.ResponseTokenExpired()
}

func carDelete(c *fiber.Ctx) error {

	level, res := services.TokenCheck(c)

	if level < 2 {
		return models.ResponseAccessDenied()
	}

	if res.Error() == models.ResponseTokenGood().Error() {
		return services.SendcarDelete(c)
	}

	return models.ResponseTokenExpired()
}
