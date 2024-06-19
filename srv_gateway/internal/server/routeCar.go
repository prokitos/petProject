package server

import (
	"github.com/gofiber/fiber/v2"
)

func tokenAccess(c *fiber.Ctx) error {

	// обращаемся к сервису авторизации для проверки токена
	// в итоге нам не надо заново вводить логин и пароль

	return c.SendStatus(fiber.StatusAccepted)
}

func tokenRefresh(c *fiber.Ctx) error {

	// если у нас истёк аксес токен, мы снова обращаемся к сервису авторизации
	// он проверяет что рефреш токен валиден, и выдаёт новый аксес токен

	return c.SendStatus(fiber.StatusAccepted)
}
