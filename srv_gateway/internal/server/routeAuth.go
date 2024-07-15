package server

import (
	"module/internal/services"

	"github.com/gofiber/fiber/v2"
)

func loginRoute(c *fiber.Ctx) error {

	return services.Authorization(c)
	// return c.SendStatus(fiber.StatusAccepted)
}

func registerRoute(c *fiber.Ctx) error {

	return services.Register(c)
	// return c.SendStatus(fiber.StatusAccepted)
}

func tokenAccess(c *fiber.Ctx) error {

	return services.TokenCheck(c)
}

func tokenRefresh(c *fiber.Ctx) error {

	return services.TokenRefresher(c)
}
