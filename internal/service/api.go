package service

import (
	"github.com/gofiber/fiber/v2"
)

type ItemService interface {
	GetItems(c *fiber.Ctx) error
	GetItem(c *fiber.Ctx) error
	CreateItem(c *fiber.Ctx) error
	UpdateItem(c *fiber.Ctx) error
	DeleteItem(c *fiber.Ctx) error
}

type UserService interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	ChangePassword(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}