package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/radekkrejcirik01/charge-points-distribution/pkg/rest/controller"
)

// Create new REST API server
func Create() *fiber.App {
	app := fiber.New()

	app.Get("/", controller.Index)

	app.Get("/get", controller.GetOutput)

	return app
}
