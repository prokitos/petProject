package services

import (
	"fmt"
	"module/internal/database"
	"module/internal/models"
)

func CarInsert(curCar models.Car) error {

	return database.CreateNewCar(curCar)

}

func CarDelete(curCar models.Car) error {

	fmt.Println("delete", curCar.Id)

	return models.ResponseGood()

}

func CarUpdate(curCar models.Car) error {

	fmt.Println("update", curCar.Id)

	return models.ResponseGood()
}

func CarShow(curCar models.Car) error {

	return database.ShowCar(curCar)
}
