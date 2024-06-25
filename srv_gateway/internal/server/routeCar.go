package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func carInsert(c *fiber.Ctx) error {

	res := tokenAccess(c)

	// тут проверка ответа. если всё хорошо, то отправлять в сервис для базы данных
	fmt.Println(res)

	return c.SendStatus(fiber.StatusAccepted)
}

func carShow(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusAccepted)
}

func carUpdate(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusAccepted)
}

func carDelete(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusAccepted)
}
