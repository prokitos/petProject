package services

import (
	"fmt"
	"module/internal/database"
	"module/internal/models"
)

func CarInsert(curCar models.Car) models.ResponseCar {

	return database.CreateNewCar(curCar)

}

func CarDelete(curCar models.Car) models.ResponseCar {

	fmt.Println("delete", curCar.Id)

	return models.ResponseCar{}

}

func CarUpdate(curCar models.Car) models.ResponseCar {

	fmt.Println("update", curCar.Id)

	return models.ResponseCar{}
}

func CarShow(curCar models.Car) models.ResponseCar {

	return database.ShowCar(curCar)
}
