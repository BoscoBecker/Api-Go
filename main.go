package main

import (
	"github.com/gofiber/fiber/v2"
	RoutesPing "github.com/your/repo/Routes"
)

func main() {
	app := fiber.New()
	app.Get("/Ping", RoutesPing.Ping)
	app.Listen(":3000")
}
