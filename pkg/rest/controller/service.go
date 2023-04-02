package controller

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/radekkrejcirik01/charge-points-distribution/pkg/database"
	"github.com/radekkrejcirik01/charge-points-distribution/pkg/service"
)

func GetOutput(c *fiber.Ctx) error {
	// Get output of all charge points with distributed currents
	output, err := service.GetOutput(database.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
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

func CreateGroup(c *fiber.Ctx) error {
	maxCurrentParam := c.Params("maxCurrent")

	// Convert string maximal current to float64
	maxCurrent, parseErr := strconv.ParseFloat(maxCurrentParam, 32)
	if parseErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: parseErr.Error(),
		})
	}

	// Create group with converted maxCurrent to float32
	if err := database.CreateGroup(database.DB, float32(maxCurrent)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "Success",
		Message: fmt.Sprintf("Group with maximal current %f created!", maxCurrent),
	})
}

func AddChargePoint(c *fiber.Ctx) error {
	groupIdParam := c.Params("groupId")

	// Convert string group id to unsigned integer
	groupId, parseErr := strconv.ParseUint(groupIdParam, 10, 32)
	if parseErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: parseErr.Error(),
		})
	}

	if err := database.AddChargePoint(database.DB, uint(groupId)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "Success",
		Message: fmt.Sprintf("Charge point added for group id %d", groupId),
	})
}

func AddChargePointConnector(c *fiber.Ctx) error {
	chPointIdParam := c.Params("chargePointId")
	chPointStatusParam := c.Params("status")

	// Convert string charge point id to unsigned integer
	chPointId, parseErr := strconv.ParseUint(chPointIdParam, 10, 32)
	if parseErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: parseErr.Error(),
		})
	}

	if err := database.AddChargePointConnector(database.DB, uint(chPointId), chPointStatusParam); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		Status: "Success",
		Message: fmt.Sprintf("Charge point connector added for charge point id %d with status %s",
			chPointId,
			chPointStatusParam,
		),
	})
}
