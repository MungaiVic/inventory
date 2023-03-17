package handlers

import (
	"inv-v2/internal/service"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(group *fiber.Group, svc service.UserService) {
	userRoutes := group.Group("/users")

	userRoutes.Get("/", func(ctx *fiber.Ctx) error {
		return service.UserService.GetUsers(svc, ctx)
	})

	userRoutes.Post("/create_user", func(ctx *fiber.Ctx) error {
		return service.UserService.Register(svc, ctx)
	})
}
