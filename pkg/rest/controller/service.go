package controller

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/radekkrejcirik01/charge-points-distribution/pkg/database"
	"github.com/radekkrejcirik01/charge-points-distribution/pkg/service"
)

func GetOutput(c *fiber.Ctx) error {
	output, err := service.GetOutput(database.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	// Pretty JSON for presentation purpose only
	data, err := json.MarshalIndent(output, "", "\t")
	if err != nil {
		return nil
	}

	return c.Status(fiber.StatusOK).Send(data)
}
