package services

import (
	"module/internal/database"
	"module/internal/models"

	log "github.com/sirupsen/logrus"
)

func CarInsert(curCar models.Car) models.ResponseCar {

	// Обагощение данных, если пусто
	if curCar.Devices == nil || curCar.Engine.EngineCapacity == 0 || curCar.Engine.EnginePower == 0 || curCar.OwnerList == nil || curCar.Mark == "" || curCar.Price == 0 {
		res, err := registerSend(curCar)
		if err != nil {
			log.Error("enricht error")
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
		if curCar.Price == 0 || curCar.Mark == "" {
			curCar.Color = res.Color
			curCar.Mark = res.Mark
			curCar.MaxSpeed = res.MaxSpeed
			curCar.Price = res.Price
			curCar.SeatsNum = res.SeatsNum
			curCar.Year = res.Year
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

func CarSellInsert(curSell models.SellingToRM) models.ResponseSell {
	return database.CreateNewSell(curSell)
}

func CarSellDelete(curSell models.SellingToRM) models.ResponseSell {

	return database.DeleteSell(curSell)
}

func CarSellUpdate(curSell models.SellingToRM) models.ResponseSell {

	return database.UpdateSell(curSell)
}

func CarSellShow(curSell models.SellingToRM) models.ResponseSell {

	return database.ShowSell(curSell)
}
