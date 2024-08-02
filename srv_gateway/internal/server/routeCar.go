package server

import (
	"module/internal/models"
	"module/internal/services"

	"github.com/gofiber/fiber/v2"

	log "github.com/sirupsen/logrus"
)

func carInsert(c *fiber.Ctx) error {

	level, res := services.TokenCheck(c)

	if level < 2 {
		log.Debug("access level too low")
		return models.ResponseAccessDenied()
	}

	if res.Error() == models.ResponseTokenGood().Error() {
		return services.SendcarInsert(c)
	}
	log.Debug("token check error")

	return models.ResponseTokenExpired()
}

func carShow(c *fiber.Ctx) error {

	level, res := services.TokenCheck(c)

	if level < 1 {
		log.Debug("access level too low")
		return models.ResponseAccessDenied()
	}

	if res.Error() == models.ResponseTokenGood().Error() {
		return services.SendcarShow(c)
	}
	log.Debug("token check error")

	return models.ResponseTokenExpired()
}

func carUpdate(c *fiber.Ctx) error {

	level, res := services.TokenCheck(c)

	if level < 2 {
		log.Debug("access level too low")
		return models.ResponseAccessDenied()
	}

	if res.Error() == models.ResponseTokenGood().Error() {
		return services.SendcarUpdate(c)
	}
	log.Debug("token check error")

	return models.ResponseTokenExpired()
}

func carDelete(c *fiber.Ctx) error {

	level, res := services.TokenCheck(c)

	if level < 2 {
		log.Debug("access level too low")
		return models.ResponseAccessDenied()
	}

	if res.Error() == models.ResponseTokenGood().Error() {
		return services.SendcarDelete(c)
	}
	log.Debug("token check error")

	return models.ResponseTokenExpired()
}
