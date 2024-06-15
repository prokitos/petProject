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
		return models.Users{}, c.SendStatus(fiber.StatusAccepted)
	}

	return curUser, nil

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
