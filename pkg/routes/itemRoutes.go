package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/MungaiVic/inventory/pkg/models"
)

func SetupItemRoutes(app *fiber.App, db *gorm.DB) {
	itemRoutes := app.Group("/api/v1")

	item := models.Item{}

	itemRoutes.Get("/items/", func(context *fiber.Ctx) error {
		return item.GetItems(context, db)
	})

	itemRoutes.Get("/getItem/:id", func(context *fiber.Ctx) error {
		return item.GetItem(context, db)
	})

	itemRoutes.Post("/items/", func(context *fiber.Ctx) error {
		return item.CreateItem(context, db)
	})
}
