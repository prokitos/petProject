package server

import (
	"log"
	"module/internal/models"

	"github.com/gofiber/fiber/v2"
)

func ServerStart(port models.ServerConfig) *fiber.App {

	app := fiber.New()

	handlers(app)

	log.Fatal(app.Listen(port.Port))

	return app
}

func handlers(instance *fiber.App) {

	instance.Post("/login", loginRoute)
	instance.Post("/register", registerRoute)
	instance.Post("/refresh", tokenRefresh)

	instance.Post("/car", carInsert)
	instance.Delete("/car", carDelete)
	instance.Get("/car", carShow)
	instance.Put("/car", carUpdate)

	// будем отправлять айди машины и айди клиента.
	// на клиенте будет идти выбор машин и клиентов по имени и названию, и отправлять айдишники в эти роуты
	instance.Post("/carSell", carSellInsert)
	instance.Delete("/carSell", carSellDelete)
	instance.Get("/carSell", carSellShow)
	instance.Put("/carSell", carSellUpdate)
}
