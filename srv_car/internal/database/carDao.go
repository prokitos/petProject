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

func DeleteCar(curModel models.Car) models.ResponseCar {

	var curCar models.Car

	result := GlobalHandler.Where(curModel).Delete(curCar)

	if result.RowsAffected == 0 || result.Error != nil {
		return models.ResponseCarBadDelete()
	}

	// Send a 201 created response
	return models.ResponseCarGoodDelete()
}

func UpdateCar(curModel models.Car) models.ResponseCar {

	id := curModel.Id
	var curCar models.Car

	if result := GlobalHandler.Preload("Engine").First(&curCar, id); result.Error != nil {
		return models.ResponseCarBadUpdate()
	}

	curCar = curModel

	GlobalHandler.Save(&curCar)

	//if GlobalHandler.Model(&curCar).Where("id = ?", id).Updates(&curModel).RowsAffected == 0 {

	// if gorm.IsRecordNotFoundError(err){
	// 	db.Create(&newUser)  // create new record
	// }

	// 	return models.ResponseCarBadUpdate()
	// }

	// Send a 201 created response
	return models.ResponseCarGoodUpdate()
}
