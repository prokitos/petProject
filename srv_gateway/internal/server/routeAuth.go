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

func tokenAccess(c *fiber.Ctx) error {

	return services.TokenCheck(c)
}

func tokenRefresh(c *fiber.Ctx) error {

	// если у нас истёк аксес токен, мы снова обращаемся к сервису авторизации
	// он проверяет что рефреш токен валиден, и выдаёт новый аксес токен

	return c.SendStatus(fiber.StatusAccepted)
}
