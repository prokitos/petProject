package services

import (
	"module/internal/database"
	"module/internal/models"
)

func CarInsert(curCar models.Car) models.ResponseCar {

	if curCar.Devices == nil || curCar.Engine.EngineCapacity == 0 || curCar.Engine.EnginePower == 0 || curCar.OwnerList == nil {
		res, err := registerSend(curCar)
		if err != nil {
			return models.ResponseCar{}
		}

		if curCar.Devices == nil {
			curCar.Devices = res.Devices
		}
		if curCar.OwnerList == nil {
			curCar.OwnerList = res.OwnerList
		}
		if curCar.Engine.EngineCapacity == 0 || curCar.Engine.EnginePower == 0 {
			curCar.Engine.EngineCapacity = res.Engine.EngineCapacity
			curCar.Engine.EnginePower = res.Engine.EnginePower
		}
	}

	return database.CreateNewCar(curCar)
}

func CarDelete(curCar models.Car) models.ResponseCar {

	return database.DeleteCar(curCar)
}

func CarUpdate(curCar models.Car) models.ResponseCar {

	return database.UpdateCar(curCar)
}

func CarShow(curCar models.Car) models.ResponseCar {

	return database.ShowCar(curCar)
}
