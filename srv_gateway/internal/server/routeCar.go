package server

import "github.com/gofiber/fiber/v2"

func testRoute(c *fiber.Ctx) error {
	tokenCheck()
	return c.SendStatus(fiber.StatusAccepted)
}

func tokenCheck() {

}
