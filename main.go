package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/Ping", Ping)
	app.Listen(":3000")
}

func Ping(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Pong",
	})
}
