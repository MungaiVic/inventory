package controllers

import (
	"github.com/MungaiVic/inventory/pkg/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetUsers(context *fiber.Ctx, db *gorm.DB) error{
	var users []models.User
	db.Find(&users)
	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": users,
	})
}
