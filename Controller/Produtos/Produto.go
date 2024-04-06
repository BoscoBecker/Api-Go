package Produto

import (
	"github.com/gofiber/fiber/v2"
	Connect "github.com/your/repo/Connection"
)

func Produto(c *fiber.Ctx) error {
	Connect.ConnectSQLServer()
	resultados, err := Connect.ExecutarConsulta("SELECT * FROM produtos")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Server internal Error")
	}
	return c.JSON(resultados)
}

func ProdutoById(c *fiber.Ctx) error {
	Connect.ConnectSQLServer()
	resultados, err := Connect.ExecutarConsultaById("SELECT * FROM produtos", 1)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Server internal Error")
	}
	return c.JSON(resultados)
}
