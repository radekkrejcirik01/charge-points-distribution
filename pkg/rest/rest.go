package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/radekkrejcirik01/charge-points-distribution/pkg/rest/controller"
)

// Create new REST API server
func Create() *fiber.App {
	app := fiber.New()

	app.Get("/get", controller.GetOutput)

	// Get method to simplify adding new records on localhost
	app.Get("/create/group/:maxCurrent", controller.CreateGroup)
	app.Get("/add/charge-point/:groupId", controller.AddChargePoint)
	app.Get("/add/charge-point-connector/:chargePointId/:status",
		controller.AddChargePointConnector,
	)

	return app
}
