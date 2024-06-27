package services

import (
	"fmt"
	"module/internal/models"
)

func CarInsert(curCar models.Car) error {

	fmt.Println("insert", curCar.Mark)

	return models.ResponseGood()
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

	fmt.Println("show", curCar.Mark)

	return models.ResponseGood()
}
