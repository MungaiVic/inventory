package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/MungaiVic/inventory/pkg/controllers"
)

func SetupItemRoutes(group *fiber.Group, db *gorm.DB) {
	itemRoutes := group.Group("/items")

	itemRoutes.Get("/", func(context *fiber.Ctx) error {
		return controllers.GetItems(context, db)
	})

	itemRoutes.Get("/getItem/:id", func(context *fiber.Ctx) error {
		return controllers.GetItem(context, db)
	})

	itemRoutes.Post("/create_item/", func(context *fiber.Ctx) error {
		return controllers.CreateItem(context, db)
	})
}
