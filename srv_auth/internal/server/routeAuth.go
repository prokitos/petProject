package server

import (
	"fmt"
	"module/internal/services"

	"github.com/gofiber/fiber/v2"
)

func loginRoute(c *fiber.Ctx) error {
	res, err := services.Authorization(c)

	// вернет bad request при ошибке, и accepter если логин и пароль есть в базе
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": err, "data": res})
}

func registerRoute(c *fiber.Ctx) error {

	res, err := services.Registration(c)

	fmt.Println(err)
	// err to string?

	// вернет bad request при ошибке, и accepter если регистрация успешна
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": err, "data": res})
}
