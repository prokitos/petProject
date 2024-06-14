package server

import "github.com/gofiber/fiber/v2"

func loginRoute(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusAccepted)
}

func registerRoute(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusAccepted)
}
