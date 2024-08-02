package database

import (
	"module/internal/models"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func CreateNewUser(c *fiber.Ctx, curUser models.Users) (models.Users, error) {

	if result := GlobalHandler.Create(&curUser); result.Error != nil {
		log.Debug("create record error!")
		return models.Users{}, models.ResponseErrorAtServer()
	}

	return curUser, nil

}

func CheckUserName(c *fiber.Ctx, curUser models.Users) error {

	var finded []models.Users

	results := GlobalHandler.Find(&finded, curUser)
	if results.Error != nil {
		log.Debug("error get records from db")
		return models.ResponseErrorAtServer()
	}

	if len(finded) != 0 {
		log.Debug("nothing to show")
		return models.ResponseBadRequest()
	}

	return nil
}

func GetExistingUser(c *fiber.Ctx, curUser models.Users) (models.Users, error) {

	var finded []models.Users

	results := GlobalHandler.Find(&finded, curUser)
	if results.Error != nil {
		log.Debug("error get records from db")
		return models.Users{}, models.ResponseErrorAtServer()
	}

	if len(finded) == 0 {
		log.Debug("nothing to show")
		return models.Users{}, models.ResponseBadRequest()
	}

	return finded[0], nil
}

func UpdateRefreshToken(c *fiber.Ctx, curUser models.Users) error {

	var test models.Users
	test.Login = curUser.Login

	result := GlobalHandler.Where(test).Updates(curUser)
	if result.Error != nil {
		log.Debug("update error in database")
		return models.ResponseErrorAtServer()
	}

	return nil
}

func CheckRefreshToken(c *fiber.Ctx, refresh string, curUser models.Users) error {

	var finded []models.Users

	results := GlobalHandler.Find(&finded, curUser)
	if results.Error != nil {
		log.Debug("database find error")
		return models.ResponseErrorAtServer()
	}

	if len(finded) == 0 {
		log.Debug("nothing to show")
		return models.ResponseBadRequest()
	}

	tokenBD := finded[0].RefreshToken
	if tokenBD != refresh {
		log.Debug("token mismatch in database")
		return models.ResponseBadRequest()
	}

	return nil
}
