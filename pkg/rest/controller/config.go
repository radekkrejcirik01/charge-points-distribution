package controller

import (
	"github.com/gofiber/fiber/v2"
)

func Config(c *fiber.Ctx) error {
	host, port, user, password, dbname := database.GetConfig()
	return c.JSON(struct {
		Host     string `json:""`
		Port     string `json:""`
		User     string `json:""`
		Password string `json:""`
		DB       string `json:""`
	}{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DB:       dbname,
	})
}
