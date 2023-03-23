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
	GetUsers(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	// ChangePassword(c *fiber.Ctx) error
	// UpdateUser(c *fiber.Ctx) error
	// DeleteUser(c *fiber.Ctx) error
}

type AuthService interface {
	Login(c *fiber.Ctx) error
}

type UserRegistration struct {
	FirstName string  `json:"first_name,omitempty"`
	LastName  string  `json:"last_name,omitempty"`
	Email     *string `json:"email,omitempty"`
	Username  string  `json:"username,omitempty"`
	Password  string  `json:"password,omitempty"`
}

type UserResponse struct {
	FirstName string  `json:"first_name,omitempty"`
	LastName  string  `json:"last_name,omitempty"`
	Email     *string `json:"email,omitempty"`
	Username  string  `json:"username,omitempty"`
}

type ItemResponse struct {
	Name       string `json:"name,omitempty"`
	Price      uint16 `json:"price,omitempty"`
	Quantity   uint16 `json:"quantity,omitempty"`
	Reorderlvl uint16 `json:"reorderlvl,omitempty"`
}

type LoginCredentials struct {
	Identifier string `json:"identifier,omitempty"`
	Password   string `json:"password,omitempty"`
}
