package server

import (
	"github.com/gofiber/fiber/v2"
)

func setupRouter(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello World!",
		})
	})
}
