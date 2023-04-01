package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Welcome page of the app
func Index(c *fiber.Ctx) error {
	fmt.Fprintln(c.Response().BodyWriter(), "Welcome! 👋")
	fmt.Fprintln(c.Response().BodyWriter(), "This app creates and updates groups of charge points. The output is a list of charge points with currents values")
	fmt.Fprintln(c.Response().BodyWriter(), "Paths 👇")
	fmt.Fprintln(c.Response().BodyWriter(), "/get - get output of all groups and their charge points with it's currents")

	return nil
}
