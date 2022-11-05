package engine

import (
	"github.com/alwindoss/poros"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Run runs the engine
func Run(cfg *poros.Config) error {
	app := fiber.New()

	app.Use(cors.New())

	// Setup the routes
	setupFiberRoutes(app, cfg)

	err := app.Listen(":3000")
	return err
}
