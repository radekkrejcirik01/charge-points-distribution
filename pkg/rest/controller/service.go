package controller

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/radekkrejcirik01/charge-points-distribution/pkg/database"
	"github.com/radekkrejcirik01/charge-points-distribution/pkg/service"
)

// Get all charge points and currents from charge points groups
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

// Create new record of group with maxiamal current in database
func CreateGroup(c *fiber.Ctx) error {
	t := &database.Group{}

	// Parse request body
	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	// Create group with converted maxCurrent to float32
	if err := database.CreateGroup(database.DB, t.MaxCurrent); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  "Success",
		Message: fmt.Sprintf("Group with maximal current %f created!", t.MaxCurrent),
	})
}

// Add new charge point with group id to database
func AddChargePoint(c *fiber.Ctx) error {
	t := &database.ChargePoint{}

	// Parse request body
	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	if err := database.AddChargePoint(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		Status: "Success",
		Message: fmt.Sprintf("Charge point added for group id %d with priority %d",
			t.GroupId,
			t.Priority,
		),
	})
}

// Add charge point connector with charge point id and status to database
func AddChargePointConnector(c *fiber.Ctx) error {
	t := &database.ChargePointConnector{}

	// Parse request body
	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	if err := database.AddChargePointConnector(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		Status: "Success",
		Message: fmt.Sprintf("Charge point connector added for charge point id %d with status %s",
			t.ChargePointId,
			t.Status,
		),
	})
}

// Update charge point connector status by id and charge point id in database
func UpdateChargePointConnector(c *fiber.Ctx) error {
	t := &database.ChargePointConnector{}

	// Parse request body
	if err := c.BodyParser(t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	if err := database.UpdateChargePointConnector(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(Response{
		Status: "Success",
		Message: fmt.Sprintf("Charge point connector with id %d updated with status %s",
			t.ChargePointId,
			t.Status,
		),
	})
}
