package services

import (
	"errors"
	"module/internal/models"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {

	var sendData models.TokenResponser
	sendData.Login = c.Query("login", "")
	sendData.Password = c.Query("password", "")
	sendData.PasswordConfirm = c.Query("password_confirm", "")

	res, err := sendToSecond(c, sendData, "register")
	if err.Error() != errors.New("good").Error() {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": err.Error(), "accessToken": res.AccessToken, "refreshToken": res.RefreshToken})
}

func Authorization(c *fiber.Ctx) error {

	var sendData models.TokenResponser
	sendData.Login = c.Query("login", "")
	sendData.Password = c.Query("password", "")

	res, err := sendToSecond(c, sendData, "login")
	if err.Error() != errors.New("good").Error() {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": err.Error(), "accessToken": res.AccessToken})
}
