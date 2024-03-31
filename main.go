package main

import (
	"github.com/gofiber/fiber/v2"
	RoutesPing "github.com/your/repo/Controller/Ping"
	Produto "github.com/your/repo/Controller/Produtos"
)

func main() {
	app := fiber.New()
	app.Get("/Ping", RoutesPing.Ping)
	app.Get("/produto/", Produto.Produto)
	app.Get("/produto/:id", Produto.ProdutoById)
	app.Listen(":3000")
}
