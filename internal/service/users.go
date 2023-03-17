package service

import (
	"inv-v2/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type UserImpl struct {
	db repository.UserRepository
}

func NewUserService(db repository.UserRepository) UserService {
	return &UserImpl{db: db}
}

func (user UserImpl) GetUsers(c *fiber.Ctx) error {
	users, _ := user.db.GetAllUsers()
	return c.Status(fiber.StatusOK).JSON(users)
}

func (user UserImpl) Register(c *fiber.Ctx) error {
	panic("Implement me!")
}
