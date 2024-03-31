package Produto

import (
	"github.com/gofiber/fiber/v2"
)

func Produto(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Produto",
	})
}
func ProdutoById(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Produto " + c.Params("id"),
	})
}
