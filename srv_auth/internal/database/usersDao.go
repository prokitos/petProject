package database

import (
	"errors"
	"module/internal/models"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func CreateNewUser(c *fiber.Ctx, curUser models.Users) (models.Users, error) {

	if result := GlobalHandler.Create(&curUser); result.Error != nil {
		log.Debug("create record error!")
		return models.Users{}, errors.New("custom error")
	}

	return curUser, nil

}

func CheckUserName(c *fiber.Ctx, curUser models.Users) error {

	var finded []models.Users

	results := GlobalHandler.Find(&finded, curUser)
	if results.Error != nil {
		return errors.New("custom error")
	}

	if len(finded) != 0 {
		return errors.New("custom error")
	}

	return nil
}

func GetExistingUser(c *fiber.Ctx, curUser models.Users) (models.Users, error) {

	var finded []models.Users

	results := GlobalHandler.Find(&finded, curUser)
	if results.Error != nil {
		return models.Users{}, errors.New("custom error")
	}

	if len(finded) == 0 {
		return models.Users{}, errors.New("custom error")
	}

	return finded[0], nil
}

func UpdateRefreshToken(c *fiber.Ctx, curUser models.Users) error {

	var test models.Users
	test.Login = curUser.Login

	//result := GlobalHandler.Model(&test).Where(test).Updates(curUser)
	result := GlobalHandler.Where(test).Updates(curUser)
	if result.Error != nil {
		return errors.New("custom error")
	}

	return nil
}
