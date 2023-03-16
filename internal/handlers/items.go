package handlers

import (
	"github.com/gofiber/fiber/v2"
	"inv-v2/internal/service"
)

func SetupItemRoutes(group *fiber.Group, svc service.ItemService) {
	itemRoutes := group.Group("/items")

	itemRoutes.Get("/", func(ctx *fiber.Ctx) error {
		return service.ItemService.GetItems(svc, ctx)
	})
	itemRoutes.Get("/:id", func(ctx *fiber.Ctx) error {
		return service.ItemService.GetItem(svc, ctx)
	})
	itemRoutes.Post("/create_item", func(ctx *fiber.Ctx) error {
		return service.ItemService.CreateItem(svc, ctx)
	})
	itemRoutes.Patch("/update_item", func(ctx *fiber.Ctx) error {
		return service.ItemService.UpdateItem(svc, ctx)
	})
	itemRoutes.Delete("/delete_item/:id", func(ctx *fiber.Ctx) error {
		return service.ItemService.DeleteItem(svc, ctx)
	})
}
