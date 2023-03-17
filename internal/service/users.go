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
	return c.Status(fiber.StatusOK).JSON(ConvertUserModelsToUserResponses(users))
}

func (user UserImpl) Register(c *fiber.Ctx) error {
	userReg := &UserRegistration{}
	err := c.BodyParser(userReg)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
			"data":    err,
		})
	}
	if err := ValidateRegisterUser(userReg); len(err) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
			"errors":  err,
		})
	}
	// TODO: Check if user exists
	hashedP, err := HashPassword(userReg.Password)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(err)
	}
	userReg.Password = hashedP
	userObj, err := user.db.CreateUser(ConvertUserRegToUserModel(*userReg))
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}
	userResp := ConvertUserRegToUserResponse(*userObj)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user successfully created",
		"data":    userResp,
	})
}
