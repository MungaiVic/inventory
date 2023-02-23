package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/MungaiVic/inventory/pkg/controllers"
)

func SetupItemRoutes(app *fiber.App, db *gorm.DB) {
	itemRoutes := app.Group("/api/v1")

	// item := models.Item{}

	itemRoutes.Get("/items/", func(context *fiber.Ctx) error {
		return controllers.GetItems(context, db)
	})

	itemRoutes.Get("/getItem/:id", func(context *fiber.Ctx) error {
		return controllers.GetItem(context, db)
	})

	itemRoutes.Post("/items/", func(context *fiber.Ctx) error {
		return controllers.CreateItem(context, db)
	})
}
