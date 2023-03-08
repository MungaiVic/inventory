package routes

import (
	"github.com/gofiber/fiber/v2"
	// "gorm.io/gorm"

	"github.com/MungaiVic/inventory/pkg/controllers"
	"github.com/MungaiVic/inventory/pkg/svc"
)

func SetupItemRoutes(group *fiber.Group, svc *svc.SVC) {
	itemRoutes := group.Group("/items")

	itemRoutes.Get("/", func(context *fiber.Ctx) error {
		return controllers.GetItems(context, svc)
	})

	// itemRoutes.Get("/get_item/:id", func(context *fiber.Ctx) error {
	// 	return controllers.GetItem(context, db)
	// })

	// itemRoutes.Post("/create_item/", func(context *fiber.Ctx) error {
	// 	return controllers.CreateItem(context, db)
	// })

	// itemRoutes.Patch("/update_item/", func(context *fiber.Ctx) error {
	// 	return controllers.UpdateItem(context, db)
	// })

	// itemRoutes.Delete("/delete_item/:id", func(context *fiber.Ctx) error {
	// 	return controllers.DeleteItem(context, db)
	// }
// )
}
