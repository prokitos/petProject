package server

import (
	"module/internal/models"
	"module/internal/services"

	"github.com/gofiber/fiber/v2"
)

func carSellInsert(c *fiber.Ctx) error {

	level, res := services.TokenCheck(c)

	if level < 3 {
		return models.ResponseAccessDenied()
	}

	if res.Error() == models.ResponseTokenGood().Error() {
		return services.SendSellcarInsert(c)
	}

	return models.ResponseTokenExpired()
}

func carSellShow(c *fiber.Ctx) error {

	level, res := services.TokenCheck(c)

	if level < 2 {
		return models.ResponseAccessDenied()
	}

	if res.Error() == models.ResponseTokenGood().Error() {
		return services.SendSellcarShow(c)
	}

	return models.ResponseTokenExpired()
}

func carSellUpdate(c *fiber.Ctx) error {

	level, res := services.TokenCheck(c)

	if level < 3 {
		return models.ResponseAccessDenied()
	}

	if res.Error() == models.ResponseTokenGood().Error() {
		return services.SendSellcarUpdate(c)
	}

	return models.ResponseTokenExpired()
}

func carSellDelete(c *fiber.Ctx) error {

	level, res := services.TokenCheck(c)

	if level < 3 {
		return models.ResponseAccessDenied()
	}

	if res.Error() == models.ResponseTokenGood().Error() {
		return services.SendSellcarDelete(c)
	}

	return models.ResponseTokenExpired()
}
