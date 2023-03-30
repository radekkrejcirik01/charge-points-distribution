package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/radekkrejcirik01/charge-points-distribution/pkg/database"
	"github.com/radekkrejcirik01/charge-points-distribution/pkg/model/statuses"
)

func GetStatusById(c *fiber.Ctx) error {
	id := c.Params("id")

	t := &statuses.Statuses{}

	if err := statuses.GetStatusById(database.DB, t, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(ResponseGetStatusById{
		Status: "Success",
		Data:   *t,
	})
}
