package services

import (
	"context"
	"module/internal/config"
	"module/internal/generpc"
	"module/internal/models"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func registerSend(car models.Car) (*models.Car, error) {

	ctxTime := time.Second * 2
	ctx, cancel := context.WithTimeout(context.Background(), ctxTime)
	defer cancel()

	conn, err := grpc.Dial(config.ExternalAddress.EnrichtService, grpc.WithInsecure())
	if err != nil {
		log.Error("grpc connect error")
		return nil, models.ResponseConnectionError()
	}
	defer conn.Close()
	client := generpc.NewEnrichmentClient(conn)

	var sendedData generpc.CarRequest
	sendedData.Mark = car.Mark
	sendedData.Year = car.Year
	sendedData.Price = int64(car.Price)
	sendedData.Color = car.Color
	sendedData.MaxSpeed = int64(car.MaxSpeed)
	sendedData.SeatsNum = int64(car.SeatsNum)
	sendedData.Status = car.Status

	response, err := client.CarEnricht(ctx, &sendedData)
	if err != nil {
		log.Error("request timeout, more than ", ctxTime)
		return nil, models.ResponseErrorAtServer()
	}

	var newCar models.Car
	newCar.Engine.EngineCapacity = float64(response.Engine.EngineCapacity)
	newCar.Engine.EnginePower = float64(response.Engine.EnginePower)

	for i := 0; i < len(response.Devices); i++ {
		var curDev models.AdditionalDevices
		curDev.DeviceName = response.Devices[i].DeviceName
		newCar.Devices = append(newCar.Devices, curDev)
	}

	for i := 0; i < len(response.OwnerList); i++ {
		var curOwn models.People
		curOwn.Name = response.OwnerList[i].Name
		curOwn.Surname = response.OwnerList[i].Surname
		curOwn.Email = response.OwnerList[i].Email
		newCar.OwnerList = append(newCar.OwnerList, curOwn)
	}

	newCar.Color = response.Color
	newCar.Mark = response.Mark
	newCar.MaxSpeed = int(response.MaxSpeed)
	newCar.Price = int(response.Price)
	newCar.SeatsNum = int(response.SeatsNum)
	newCar.Year = response.Year

	return &newCar, nil
}
