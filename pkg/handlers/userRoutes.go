package routes

import (
	"github.com/MungaiVic/inventory/pkg/controllers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupUserRoutes(group *fiber.Group, db *gorm.DB){
	userRoutes := group.Group("/users")

	userRoutes.Get("/", func(context *fiber.Ctx) error {
		return controllers.GetUsers(context, db)
	})
}
