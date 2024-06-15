package server

import (
	"fmt"
	"module/internal/services"

	"github.com/gofiber/fiber/v2"
)

func loginRoute(c *fiber.Ctx) error {
	res, err := services.Authorization(c)
	fmt.Println(res)

	// вернет bad request при ошибке, и accepter если логин и пароль есть в базе
	return err
}

func registerRoute(c *fiber.Ctx) error {

	res, err := services.Registration(c)
	fmt.Println(res)
	return err
}
