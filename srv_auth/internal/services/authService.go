package services

import (
	"module/internal/database"
	"module/internal/models"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func Authorization(c *fiber.Ctx) (models.TokenResponser, error) {

	var curUser models.Users
	curUser.Login = c.Query("login", "")
	curUser.Password = c.Query("password", "")

	if len(curUser.Login) < 5 || len(curUser.Password) < 5 {
		log.Debug("too short login or password")
		return models.TokenResponser{}, models.ResponseBadRequest()
	}

	// проверка пользователя и его уровня доступа. потом генерация токена
	curUser, err := database.GetExistingUser(c, curUser)
	if err != nil {
		log.Debug("database dont have current user")
		return models.TokenResponser{}, models.ResponseBadRequest()
	}

	res := TokenGetPair(curUser)
	curUser.RefreshToken = res.RefreshToken

	// записываем новый рефреш токен в базу
	err = database.UpdateRefreshToken(c, curUser)
	if err != nil {
		log.Debug("error add refresh token in db")
		return models.TokenResponser{}, models.ResponseTokenError()
	}

	return res, models.ResponseGood()
}

func Registration(c *fiber.Ctx) (models.TokenResponser, error) {

	login := c.Query("login", "")
	password := c.Query("password", "")
	confirmPassword := c.Query("password_confirm", "")

	if password != confirmPassword {
		log.Debug("password and confirm password dont same")
		return models.TokenResponser{}, models.ResponseBadRequest()
	}
	if len(login) < 5 || len(password) < 5 {
		log.Debug("too short login or password")
		return models.TokenResponser{}, models.ResponseBadRequest()
	}

	var curUser models.Users
	curUser.Login = login
	curUser.Password = password
	curUser.AccessLevel = 1

	res := TokenGetPair(curUser)
	curUser.RefreshToken = res.RefreshToken

	// проверка что нет пользователя с таким именем
	err := database.CheckUserName(c, models.Users{Login: login})
	if err != nil {
		log.Debug("database already have current login")
		return models.TokenResponser{}, models.ResponseUserExist()
	}

	// создание нового пользователя с 1 уровнем доступа. потом генерация токена
	curUser, err = database.CreateNewUser(c, curUser)
	if err != nil {
		log.Debug("database error at create new user")
		return models.TokenResponser{}, models.ResponseTokenError()
	}

	return res, models.ResponseGood()
}

func RefreshTokenNew(c *fiber.Ctx, refresh string, access string) error {

	// проверка что рефреш токен не истёк, также проверка что аксес токен совпадает с рефреш токеном, и вывод логина пользователя из токена.
	user, err := GetUserFromRefresh(refresh, access)
	if err != nil {
		log.Debug("error get user from refresh token")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": err.Error()})
	}

	// потом проверка что в базе данных лежит наш рефреш токен
	err = database.CheckRefreshToken(c, refresh, user)
	if err != nil {
		log.Debug("error get token from database")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": err.Error()})
	}

	// и уже потом генерация нового рефреш и аксес токена
	token := TokenGetPair(user)

	// записываем токен в базу
	user.RefreshToken = token.RefreshToken
	err = database.UpdateRefreshToken(c, user)
	if err != nil {
		log.Debug("error write token to db")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": models.ResponseGood().Error(), "data": token})
}

func DebugUpgrade(c *fiber.Ctx) error {

	var curUser models.Users
	curUser.Login = c.Query("login", "")

	// добавление уровня к пользователю
	err := database.UpdateUserLevel(c, curUser)
	if err != nil {
		log.Debug("error add level")
		return models.ResponseTokenError()
	}

	return models.ResponseGood()
}
