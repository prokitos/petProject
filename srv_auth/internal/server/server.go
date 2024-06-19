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

	instance.Post("/accessToken", checkAccessToken)
	instance.Post("/refreshToken", checkRefreshToken)

}
