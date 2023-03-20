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

func (user UserImpl) GetUser(c *fiber.Ctx) error {
	emailParam := c.Query("email")
	usernameParam := c.Query("username")
	userID := c.Query("user_id")

	switch {
	case emailParam != "":
		userObj, err := user.db.GetUserByEmail(emailParam)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(ConvertUserModelToUserResponse(userObj))
	case usernameParam != "":
		userObj, err := user.db.GetUserByUsername(usernameParam)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(ConvertUserModelToUserResponse(userObj))
	case userID != "":
		userObj, err := user.db.GetUserByID(userID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(ConvertUserModelToUserResponse(userObj))
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Please supply email, username or user_id",
		})
	}

}

func (user UserImpl) Register(c *fiber.Ctx) error {
	userReg := &UserRegistration{}
	err := c.BodyParser(userReg)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
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
	userExists, err := user.db.GetUserByUsername(userReg.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	} else if userExists.Username != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User already exists",
		})
	}
	hashedP, err := HashPassword(userReg.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	userReg.Password = hashedP
	userObj, err := user.db.CreateUser(ConvertUserRegToUserModel(*userReg))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}
	userResp := ConvertUserRegToUserResponse(*userObj)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user successfully created",
		"data":    userResp,
	})
}
