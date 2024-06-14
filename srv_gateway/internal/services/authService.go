package services

import "github.com/gofiber/fiber/v2"

func register(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusAccepted)
}

func authorization(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusAccepted)
}
