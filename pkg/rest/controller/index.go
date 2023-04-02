package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Welcome page of the app
func Index(c *fiber.Ctx) error {
	fmt.Fprintln(c.Response().BodyWriter(), "Welcome! 👋")
	fmt.Fprintln(c.Response().BodyWriter(), "This application distributes maximum current of charge group between charge points based on charging status. Full docs can be found at repository README or docs subfolder")
	fmt.Fprintln(c.Response().BodyWriter(), "Get started with path /get")

	return nil
}
