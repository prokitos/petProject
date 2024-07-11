package server

import (
	"context"
	"fmt"
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

	fmt.Println(req)
	fmt.Println("пришло")

	test1 := services.EnrichtedOwner()
	test2 := services.EnrichtedEngine()
	test3 := services.EnrichtedDevices()

	fmt.Println(test1)
	fmt.Println(test2)
	fmt.Println(test3)

	return ResultCompile(test1, test2, test3), nil

}

func ResultCompile(test1 []models.People, test2 models.CarEngine, test3 []models.AdditionalDevices) *generpc.CarResponse {

	var engines generpc.CarEngine
	engines.EnginePower = float32(test2.EnginePower)
	engines.EngineCapacity = float32(test2.EngineCapacity)

	// var peoples []generpc.People
	// for i := 0; i < len(test1); i++ {
	// 	var curPeople generpc.People
	// 	curPeople.Name = test1[i].Name
	// 	curPeople.Surname = test1[i].Surname
	// 	curPeople.Email = test1[i].Email
	// 	peoples = append(peoples, curPeople)
	// }

	// var devices []generpc.AdditionalDevices
	// for i := 0; i < len(test3); i++ {
	// 	var curDev generpc.AdditionalDevices
	// 	curDev.DeviceName = test3[i].DeviceName
	// 	devices = append(devices, curDev)
	// }

	var result *generpc.CarResponse
	result.Engine = &engines
	// result.OwnerList = &peoples
	// result.Devices = &devices

	return result
}
