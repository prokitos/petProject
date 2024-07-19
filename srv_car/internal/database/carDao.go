package database

import (
	"module/internal/models"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
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

	// остаются не связанные записи в people. удалять или нет?
	// если удалять, то нужно получить сначала всех овнеров у машины, а потом в цикле удалять из people людей с нужным номером.
	GlobalHandler.Select(clause.Associations).Delete(&curModel)
	GlobalHandler.Delete(&models.Car{}, curModel.Id)

	// Send a 201 created response
	return models.ResponseCarGoodDelete()
}

func UpdateCar(curModel models.Car) models.ResponseCar {

	id := curModel.Id

	var curCar models.Car
	if result := GlobalHandler.Preload("Engine").Preload("Devices").Preload("OwnerList").First(&curCar, id); result.Error != nil {
		return models.ResponseCarBadUpdate()
	}

	GlobalHandler.Model(models.CarEngine{}).Where("id = ?", curCar.Engine.Id).Updates(&curModel.Engine)

	for i := 0; i < len(curModel.Devices); i++ {
		GlobalHandler.Model(models.AdditionalDevices{}).Where("id = ?", curCar.Devices[i].Id).Updates(&curModel.Devices[i])
	}

	for i := 0; i < len(curModel.OwnerList); i++ {
		GlobalHandler.Model(models.People{}).Where("id = ?", curCar.OwnerList[i].Id).Updates(&curModel.OwnerList[i])
	}

	GlobalHandler.Model(models.Car{}).Where("id = ?", curCar.Id).Updates(&curModel)

	//if GlobalHandler.Model(&curCar).Where("id = ?", id).Updates(&curModel).RowsAffected == 0 {
	// if gorm.IsRecordNotFoundError(err){
	// 	db.Create(&newUser)  // create new record
	// }

	// Send a 201 created response
	return models.ResponseCarGoodUpdate()
}

func CreateNewSell(instance models.SellingToRM) models.ResponseSell {

	var curSell models.Selling
	var curCar models.Car
	var curPeople models.People
	curCar.Id = instance.CarId
	curPeople.Id = instance.PeopleId
	curSell.Car = curCar
	curSell.People = curPeople

	if result := GlobalHandler.Create(&curSell); result.Error != nil {
		log.Debug("create record error!")
		return models.ResponseSellBadExecute()
	}

	return models.ResponseSellGoodExecute()
}

func DeleteSell(curModel models.SellingToRM) models.ResponseSell {

	GlobalHandler.Select(clause.Associations).Delete(&curModel)
	GlobalHandler.Delete(&models.Car{}, curModel.Id)

	return models.ResponseSellGoodExecute()
}

func ShowSell(instance models.SellingToRM) models.ResponseSell {

	var curSell models.Selling
	var curCar models.Car
	var curPeople models.People
	curCar.Id = instance.CarId
	curPeople.Id = instance.PeopleId
	curSell.Car = curCar
	curSell.People = curPeople

	var finded []models.Selling

	results := GlobalHandler.Preload("Car").Preload("People").Find(&finded, curSell)
	if results.Error != nil {
		return models.ResponseSellBadShow()
	}

	return models.ResponseSellGoodShow(finded)
}

func UpdateSell(curModel models.SellingToRM) models.ResponseSell {

	return models.ResponseSellBadExecute()
}
