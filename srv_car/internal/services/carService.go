package services

import (
	"module/internal/database"
	"module/internal/models"
)

func CarInsert(curCar models.Car) models.ResponseCar {

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
