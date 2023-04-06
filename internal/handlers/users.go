package handlers

import (
	"inv-v2/internal/middleware"
	"inv-v2/internal/service"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(group *fiber.Group, svc service.UserService) {
	userRoutes := group.Group("/users")

	userRoutes.Get("/", middleware.Protected(), func(ctx *fiber.Ctx) error {
		return service.UserService.GetUsers(svc, ctx)
	})

	userRoutes.Get("/get_user", middleware.AdminOnly(), middleware.Protected(), func(ctx *fiber.Ctx) error {
		return service.UserService.GetUser(svc, ctx)
	})

	userRoutes.Post("/create_user", middleware.Protected(), func(ctx *fiber.Ctx) error {
		return service.UserService.Register(svc, ctx)
	})

	userRoutes.Put("/update_user", middleware.Protected(), func(ctx *fiber.Ctx) error {
		return service.UserService.UpdateUser(svc, ctx)
	})

	userRoutes.Put("/changepass", func(ctx *fiber.Ctx) error {
		return service.UserService.ChangePassword(svc, ctx)
	})
}
