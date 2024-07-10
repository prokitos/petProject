package app

import (
	"fmt"
	"module/internal/models"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

type App struct {
	GRPCserver *grpc.Server
	Port       int
}

// подгрузка настроек в конфиг сервера
func (a *App) NewServer(port models.ServerConfig) {

	gRPCserver := grpc.NewServer()
	a.GRPCserver = gRPCserver
	innerPort, err := strconv.Atoi(port.Port)
	if err != nil {
		panic(err)
	}
	a.Port = innerPort

}

// попытка безопасно выключить сервер
func (a *App) Stop() {

	a.GRPCserver.GracefulStop()

}

// запуск сервера, он пытается запустить его, и при ошибке закрывает приложение
func (a *App) Run() {

	if err := a.StartServer(); err != nil {
		panic(err)
	}

}

func (a *App) StartServer() error {

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.Port))
	if err != nil {
		return fmt.Errorf("%s: %w", "run", err)
	}

	if err := a.GRPCserver.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", "run", err)
	}

	return nil

}
