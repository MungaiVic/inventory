package handlers

import (
	"inv-v2/internal/service"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes (group *fiber.Group, svc service.AuthService){
	authRoutes := group.Group("/auth")

	authRoutes.Post("/login", func(ctx *fiber.Ctx) error {
		return service.AuthService.Login(svc, ctx)
	})
}