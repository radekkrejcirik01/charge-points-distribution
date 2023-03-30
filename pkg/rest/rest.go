package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/radekkrejcirik01/charge-points-distribution/pkg/database"
	"github.com/radekkrejcirik01/charge-points-distribution/pkg/model/statuses"
	"github.com/radekkrejcirik01/charge-points-distribution/pkg/rest/controller"
	"gorm.io/gorm"
)

// Create new REST API server
func Create() *fiber.App {
	app := fiber.New()

	app.Get("/", controller.Index)
	app.Get("/config", controller.Config)

	app.Get("/record", GetRecord)

	return app
}

func GetRecord(c *fiber.Ctx) error {
	t := &statuses.Statuses{}
	if err := Read(database.DB, t); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(controller.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(controller.Response{
		Status: "ok",
		Data:   &t,
	})
}

func Read(db *gorm.DB, status string) error {
	return db.Table("statuses").Select("connector_status").Where("id = 2").First(&status).Error
}
