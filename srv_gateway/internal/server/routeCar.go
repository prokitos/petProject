package server

import (
	"errors"
	"module/internal/services"

	"github.com/gofiber/fiber/v2"
)

func carInsert(c *fiber.Ctx) error {

	res := tokenAccess(c)

	if res.Error() == errors.New("token useful").Error() {
		services.SendcarInsert(c)
		return c.SendStatus(fiber.StatusAccepted)
	}

	return c.SendStatus(fiber.StatusBadRequest)
}

func carShow(c *fiber.Ctx) error {

	res := tokenAccess(c)

	if res.Error() == errors.New("token useful").Error() {
		services.SendcarShow(c)
		return c.SendStatus(fiber.StatusAccepted)
	}

	return c.SendStatus(fiber.StatusBadRequest)
}

func carUpdate(c *fiber.Ctx) error {

	res := tokenAccess(c)

	if res.Error() == errors.New("token useful").Error() {
		services.SendcarUpdate(c)
		return c.SendStatus(fiber.StatusAccepted)
	}

	return c.SendStatus(fiber.StatusBadRequest)
}

func carDelete(c *fiber.Ctx) error {

	res := tokenAccess(c)

	if res.Error() == errors.New("token useful").Error() {
		services.SendcarDelete(c)
		return c.SendStatus(fiber.StatusAccepted)
	}

	return c.SendStatus(fiber.StatusBadRequest)
}
