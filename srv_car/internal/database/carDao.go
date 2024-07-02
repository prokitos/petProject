package database

import (
	"errors"
	"fmt"
	"module/internal/models"

	log "github.com/sirupsen/logrus"
)

func CreateNewCar(curUser models.Car) error {

	if result := GlobalHandler.Create(&curUser); result.Error != nil {
		log.Debug("create record error!")
		return models.ResponseErrorAtServer()
	}

	log.Debug("record created")
	return models.ResponseGood()
}

func ShowCar(curModel models.Car) error {

	var finded []models.Car

	results := GlobalHandler.Preload("Engine").Preload("Devices").Preload("OwnerList").Find(&finded, curModel)
	if results.Error != nil {
		return errors.New("FIGNYA dsgsd dsgdsgds")
	}

	fmt.Println(finded)
	return errors.New("OFIGENNO dsgsdg sdsgsd g")
}
