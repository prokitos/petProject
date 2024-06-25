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

	res, err := sendToAuth(c, sendData, "/register")
	if err.Error() != errors.New("good").Error() {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": err.Error(), "accessToken": res.AccessToken, "refreshToken": res.RefreshToken})
}

func Authorization(c *fiber.Ctx) error {

	var sendData models.TokenResponser
	sendData.Login = c.Query("login", "")
	sendData.Password = c.Query("password", "")

	res, err := sendToAuth(c, sendData, "/login")
	if err.Error() != errors.New("good").Error() {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": err.Error(), "accessToken": res.AccessToken, "refreshToken": res.RefreshToken})
}

func TokenCheck(c *fiber.Ctx) error {

	authorization := c.Get("Authorization")
	sendTokenToCheck(c, authorization, "/accessToken")

	return c.SendStatus(fiber.StatusAccepted)
}

func TokenRefresher(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusAccepted)
}
