package services

import (
	"module/internal/models"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {

	var sendData models.TokenResponser
	sendData.Login = c.Query("login", "")
	sendData.Password = c.Query("password", "")
	sendData.PasswordConfirm = c.Query("password_confirm", "")

	sendToSecond(c, sendData, "register")

	return c.SendStatus(fiber.StatusAccepted)
}

func Authorization(c *fiber.Ctx) error {

	var sendData models.TokenResponser
	sendData.Login = c.Query("login", "")
	sendData.Password = c.Query("password", "")

	sendToSecond(c, sendData, "login")

	return c.SendStatus(fiber.StatusAccepted)
}
