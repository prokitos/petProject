package server

import (
	"module/internal/services"

	"github.com/gofiber/fiber/v2"
)

func loginRoute(c *fiber.Ctx) error {

	services.Authorization(c)
	return c.SendStatus(fiber.StatusAccepted)
}

func registerRoute(c *fiber.Ctx) error {

	services.Register(c)
	return c.SendStatus(fiber.StatusAccepted)
}
