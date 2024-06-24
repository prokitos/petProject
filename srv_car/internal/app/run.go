package app

import (
	"module/internal/server"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	Server *fiber.App
}

func (a *App) NewServer(port string) {
	a.Server = server.ServerStart(port)
}
