package server

import (
	"module/internal/services"

	"github.com/gofiber/fiber/v2"
)

func loginRoute(c *fiber.Ctx) error {
	res, err := services.Authorization(c)

	// вернет bad request при ошибке, и accepter если логин и пароль есть в базе
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": err.Error(), "data": res})
}

func registerRoute(c *fiber.Ctx) error {

	res, err := services.Registration(c)

	// вернет bad request при ошибке, и accepter если регистрация успешна
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": err.Error(), "data": res})
}

func checkAccessToken(c *fiber.Ctx) error {

	// проверка токена
	authorization := c.Get("Authorization")
	status := services.TokenAccessValidate(authorization)

	// если всё нормально, то идти дальше и обращаться к микросервису
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": status})
}

func checkRefreshToken(c *fiber.Ctx) error {

	// проверка рефреш токена и выдача нового аксес токена
	authorization := c.Get("Authorization")
	token := services.GetAccessTokenFromRefresh(authorization)

	// возвращаем новый аксес токен
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"token": token})
}
