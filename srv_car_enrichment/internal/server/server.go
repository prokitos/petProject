package server

import (
	"context"
	"module/internal/models"
	"module/internal/server/generpc"
	"module/internal/services"

	"google.golang.org/grpc"
)

type serverApi struct {
	generpc.UnimplementedEnrichmentServer
}

func Register(gRPC *grpc.Server) {
	generpc.RegisterEnrichmentServer(gRPC, &serverApi{})
}

func (s *serverApi) CarEnricht(ctx context.Context, req *generpc.CarRequest) (*generpc.CarResponse, error) {

	test1 := services.EnrichtedOwner()
	test2 := services.EnrichtedEngine()
	test3 := services.EnrichtedDevices()
	test4 := services.EnrichtedBase()

	res := ResultCompile(test1, test2, test3, test4)
	return &res, nil

}

func ResultCompile(test1 []models.People, test2 models.CarEngine, test3 []models.AdditionalDevices, test4 models.Car) generpc.CarResponse {

	var result generpc.CarResponse

	var engines generpc.CarEngine
	engines.EnginePower = float32(test2.EnginePower)
	engines.EngineCapacity = float32(test2.EngineCapacity)
	result.Engine = &engines

	for i := 0; i < len(test1); i++ {
		var curPeople generpc.People
		curPeople.Name = test1[i].Name
		curPeople.Surname = test1[i].Surname
		curPeople.Email = test1[i].Email
		result.OwnerList = append(result.OwnerList, &curPeople)
	}

	for i := 0; i < len(test3); i++ {
		var curDev generpc.AdditionalDevices
		curDev.DeviceName = test3[i].DeviceName
		result.Devices = append(result.Devices, &curDev)
	}

	result.Color = test4.Color
	result.Mark = test4.Mark
	result.Year = test4.Year
	result.Price = int64(test4.Price)
	result.MaxSpeed = int64(test4.MaxSpeed)
	result.SeatsNum = int64(test4.SeatsNum)
	result.Status = "Sale"

	return result
}
