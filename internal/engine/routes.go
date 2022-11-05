package engine

import (
	"github.com/alwindoss/poros"
	"github.com/gofiber/fiber/v2"
)

func setupFiberRoutes(app *fiber.App, cfg *poros.Config) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
