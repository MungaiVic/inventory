package handlers

import (
	"inv-v2/internal/middleware"
	"inv-v2/internal/service"

	"github.com/gofiber/fiber/v2"
)

func SetupItemRoutes(group *fiber.Group, svc service.ItemService) {
	itemRoutes := group.Group("/items")

	itemRoutes.Get("/", func(ctx *fiber.Ctx) error {
		return service.ItemService.GetItems(svc, ctx)
	})
	itemRoutes.Get("/:id", func(ctx *fiber.Ctx) error {
		return service.ItemService.GetItem(svc, ctx)
	})
	itemRoutes.Post("/inventory", middleware.Protected(), func(ctx *fiber.Ctx) error {
		return service.ItemService.CreateItem(svc, ctx)
	})
	itemRoutes.Patch("/inventory", middleware.AdminOnly(), func(ctx *fiber.Ctx) error {
		return service.ItemService.UpdateItem(svc, ctx)
	})
	itemRoutes.Delete("/inventory/:id", middleware.AdminOnly(), func(ctx *fiber.Ctx) error {
		return service.ItemService.DeleteItem(svc, ctx)
	})
}
