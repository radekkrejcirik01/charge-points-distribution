package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Welcome page of the app
func Index(c *fiber.Ctx) error {
	_, err := fmt.Fprint(c.Response().BodyWriter(),
		`
		Welcome! 👋
		This app distributes maximum current of charge group between charge points based on priority.
		Full docs can be found at repository README or docs subfolder
		Get started with path /get
		`)

	return err
}
