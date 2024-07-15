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

	res, err := sendToAuth(c, sendData, "/register")
	if err.Error() != models.ResponseGood().Error() {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": err.Error(), "accessToken": res.AccessToken, "refreshToken": res.RefreshToken})
}

func Authorization(c *fiber.Ctx) error {

	var sendData models.TokenResponser
	sendData.Login = c.Query("login", "")
	sendData.Password = c.Query("password", "")

	res, err := sendToAuth(c, sendData, "/login")
	if err.Error() != models.ResponseGood().Error() {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": err.Error(), "accessToken": res.AccessToken, "refreshToken": res.RefreshToken})
}

func TokenCheck(c *fiber.Ctx) error {

	authorization := c.Get("Authorization")
	return sendTokenToCheck(c, authorization, "/accessToken")
}

func TokenRefresher(c *fiber.Ctx) error {

	authorization := c.Get("Authorization")
	refresh := c.FormValue("refresh")

	token := sendTokenToRefresh(c, authorization, refresh, "/refreshToken")
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": token.Status, "accessToken": token.Data.AccessToken, "refreshToken": token.Data.RefreshToken})
}
