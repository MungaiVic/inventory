package service

import (
	"inv-v2/internal/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func (user UserImpl) UpdateUser(ctx *fiber.Ctx) error {
	userupdate := &UserUpdate{}
	ctx.BodyParser(userupdate)
	if updateErrors := ValidateUpdateUser(userupdate); len(updateErrors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(updateErrors)
	}
	toBeUpdatedUser, err := user.db.GetUserByID(userupdate.UserID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	toBeUpdatedUser.Username = userupdate.Username
	toBeUpdatedUser.Email = userupdate.Email
	toBeUpdatedUser.FirstName = userupdate.FirstName
	toBeUpdatedUser.LastName = userupdate.LastName
	toBeUpdatedUser.UserID = uuid.Must(uuid.Parse(userupdate.UserID))

	updatedUser, err := user.db.UpdateUser(toBeUpdatedUser)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}
	userResp := ConvertUserModelToUserResponse(updatedUser)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user successfully updated",
		"data":    userResp,
	})
}

func (user UserImpl) ChangePassword(ctx *fiber.Ctx) error {
	passChange := &PasswordChange{}
	ctx.BodyParser(passChange)
	currentPasswordHash, err := user.db.GetUserByID(passChange.UserID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User does not exist",
		})
	}
	validOldPass := ValidatePassword(currentPasswordHash.Password, passChange.OldPass)
	if validOldPass {
		newPassHash, err := HashPassword(passChange.NewPass)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err)
		}
		currentPasswordHash.Password = newPassHash
		user.db.ChangePassword(currentPasswordHash)
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Password Changed successfully",
		})
	}
	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": "wrong password entered",
	})
}

func (user UserImpl) DeleteUser(ctx *fiber.Ctx) error {
	userID := ctx.Query("user_id")
	if userID != "" {
		err := user.db.DeleteUser(userID)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(err)
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "user deleted successfully.",
		})
	}
	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": "user_id cannot be empty",
	})
}
