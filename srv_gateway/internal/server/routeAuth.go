package server

import (
	"module/internal/services"

	"github.com/gofiber/fiber/v2"
)

func loginRoute(c *fiber.Ctx) error {

	return services.Authorization(c)
}

func registerRoute(c *fiber.Ctx) error {

	return services.Register(c)
}

func tokenRefresh(c *fiber.Ctx) error {

	return services.TokenRefresher(c)
}

func debugUpgrade(c *fiber.Ctx) error {

	return services.DebugUpgrade(c)
}
