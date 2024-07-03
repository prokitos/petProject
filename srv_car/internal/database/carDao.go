package database

import (
	"module/internal/models"

	log "github.com/sirupsen/logrus"
)

func CreateNewCar(curUser models.Car) models.ResponseCar {

	if result := GlobalHandler.Create(&curUser); result.Error != nil {
		log.Debug("create record error!")
		return models.ResponseCarBadCreate()
	}

	return models.ResponseCarGoodCreate()
}

func ShowCar(curModel models.Car) models.ResponseCar {

	var finded []models.Car

	results := GlobalHandler.Preload("Engine").Preload("Devices").Preload("OwnerList").Find(&finded, curModel)
	if results.Error != nil {
		return models.ResponseCarBadShow()
	}

	return models.ResponseCarGoodShow(finded)
}
