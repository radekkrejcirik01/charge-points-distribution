package rest

import (
	"github.com/gofiber/fiber/v2"
)

// Create new REST API serveer
func Create() *fiber.App {
	app := fiber.New()

	app.Get("/", controller.Index)

	app.Get("/config", controller.Config)

	return app
}
