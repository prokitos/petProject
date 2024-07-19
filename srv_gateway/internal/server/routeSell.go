package server

import (
	"module/internal/models"
	"module/internal/services"

	"github.com/gofiber/fiber/v2"
)

func carSellInsert(c *fiber.Ctx) error {

	level, res := services.TokenCheck(c)

	if res.Error() == models.ResponseTokenGood().Error() && level > 2 {
		return services.SendSellcarInsert(c)
	}

	return models.ResponseTokenExpired()
}

func carSellShow(c *fiber.Ctx) error {

	level, res := services.TokenCheck(c)

	if res.Error() == models.ResponseTokenGood().Error() && level > 1 {
		return services.SendSellcarShow(c)
	}

	return models.ResponseTokenExpired()
}

func carSellUpdate(c *fiber.Ctx) error {

	level, res := services.TokenCheck(c)

	if res.Error() == models.ResponseTokenGood().Error() && level > 2 {
		return services.SendSellcarUpdate(c)
	}

	return models.ResponseTokenExpired()
}

func carSellDelete(c *fiber.Ctx) error {

	level, res := services.TokenCheck(c)

	if res.Error() == models.ResponseTokenGood().Error() && level > 2 {
		return services.SendSellcarDelete(c)
	}

	return models.ResponseTokenExpired()
}
