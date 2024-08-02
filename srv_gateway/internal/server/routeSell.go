package server

import (
	"module/internal/models"
	"module/internal/services"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func carSellInsert(c *fiber.Ctx) error {

	level, res := services.TokenCheck(c)

	if level < 3 {
		log.Debug("access level too low")
		return models.ResponseAccessDenied()
	}

	if res.Error() == models.ResponseTokenGood().Error() {
		return services.SendSellcarInsert(c)
	}
	log.Debug("token check error")

	return models.ResponseTokenExpired()
}

func carSellShow(c *fiber.Ctx) error {

	level, res := services.TokenCheck(c)

	if level < 2 {
		log.Debug("access level too low")
		return models.ResponseAccessDenied()
	}

	if res.Error() == models.ResponseTokenGood().Error() {
		return services.SendSellcarShow(c)
	}
	log.Debug("token check error")

	return models.ResponseTokenExpired()
}

func carSellUpdate(c *fiber.Ctx) error {

	level, res := services.TokenCheck(c)

	if level < 3 {
		log.Debug("access level too low")
		return models.ResponseAccessDenied()
	}

	if res.Error() == models.ResponseTokenGood().Error() {
		return services.SendSellcarUpdate(c)
	}
	log.Debug("token check error")

	return models.ResponseTokenExpired()
}

func carSellDelete(c *fiber.Ctx) error {

	level, res := services.TokenCheck(c)

	if level < 3 {
		log.Debug("access level too low")
		return models.ResponseAccessDenied()
	}

	if res.Error() == models.ResponseTokenGood().Error() {
		return services.SendSellcarDelete(c)
	}
	log.Debug("token check error")

	return models.ResponseTokenExpired()
}
