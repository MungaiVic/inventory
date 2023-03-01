package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetUsers(context *fiber.Ctx, db *gorm.DB) error{
	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get all users",
	})
}
