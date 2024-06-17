package services

import (
	"errors"
	"fmt"
	"module/internal/database"
	"module/internal/models"

	"github.com/gofiber/fiber/v2"
)

func Authorization(c *fiber.Ctx) (models.TokenResponser, error) {

	var curUser models.Users
	curUser.Login = c.Query("login", "")
	curUser.Password = c.Query("password", "")

	if len(curUser.Login) < 5 || len(curUser.Password) < 5 {
		return models.TokenResponser{}, c.SendStatus(fiber.StatusBadRequest)
	}

	// проверка пользователя и его уровня доступа. потом генерация токена
	curUser, err := database.GetExistingUser(c, curUser)
	if err != nil {
		return models.TokenResponser{}, c.SendStatus(fiber.StatusBadRequest)
	}

	res := TokenGetPair(curUser)

	return res, c.SendStatus(fiber.StatusAccepted)
}

func Registration(c *fiber.Ctx) (models.TokenResponser, error) {

	login := c.Query("login", "")
	password := c.Query("password", "")
	confirmPassword := c.Query("password_confirm", "")

	if password != confirmPassword {
		return models.TokenResponser{}, errors.New("custom error")
	}
	if len(login) < 5 || len(password) < 5 {
		return models.TokenResponser{}, errors.New("custom error")
	}

	var curUser models.Users
	curUser.Login = login
	curUser.Password = password
	curUser.AccessLevel = 1

	res := TokenGetPair(curUser)
	curUser.RefreshToken = res.RefreshToken

	// создание нового пользователя с 1 уровнем доступа. потом генерация токена
	curUser, err := database.CreateNewUser(c, curUser)
	if err != nil {
		fmt.Println("случилась ошибка!()")
		return models.TokenResponser{}, c.SendStatus(fiber.StatusBadRequest)
	}

	return res, c.SendStatus(fiber.StatusAccepted)
}
