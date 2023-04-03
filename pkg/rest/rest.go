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

	app.Post("/group", controller.CreateGroup)
	app.Post("/charge-point", controller.AddChargePoint)
	app.Post("/charge-point-connector", controller.AddChargePointConnector)

	app.Put("/charge-point", controller.UpdateChargePoint)
	app.Put("/charge-point-connector", controller.UpdateChargePointConnector)

	return app
}
