package service

import (
	"inv-v2/internal/repository"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginImpl struct {
	db repository.UserRepository
}

// HashPassword generates a password hash for storage in the DB.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

/*
ValidatePassword will be used to check if user has entered correct password before allowing jwt token generation.

# It will take in the hashed password from the DB and a user-supplied password  and return true if they match

Example:

	ValidatePassword(`'hashedPassword'`, `userPassword`)
*/
func ValidatePassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func NewAuthService(db repository.UserRepository) AuthService {
	return &LoginImpl{db: db}
}

func (auth LoginImpl) Login(ctx *fiber.Ctx) error {
	// login logic will come here
	creds := &LoginCredentials{}
	err := ctx.BodyParser(creds)
	if creds.Identifier == "" || creds.Password == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "identifier/password must not be empty",
		})
	}
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	if strings.Contains(creds.Identifier, "@") {
		user, err := auth.db.GetUserByEmail(creds.Identifier)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err)
		}
		hashedPassword := user.Password

		if ValidatePassword(hashedPassword, creds.Password) {
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"Message": "Authentication credentials look OK.",
			})
		} else {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Credentials do not match.",
			})
		}
	} else {
		user, err := auth.db.GetUserByUsername(creds.Identifier)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(err)
		}
		hashedPassword := user.Password

		if ValidatePassword(hashedPassword, creds.Password) {
			jwtToken := jwt.New(jwt.SigningMethodHS256)
			claims := jwtToken.Claims.(jwt.MapClaims)
			claims["identity"] = user.UserID
			claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
			claims["admin"] = user.IsAdmin

			tok, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET"))) // This should be in the env file
			if err != nil{
				return ctx.SendStatus(fiber.StatusInternalServerError)
			}
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"Message": "Authentication credentials look OK.",
				"jwtToken": tok,
			})
		} else {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Credentials do not match.",
			})
		}
	}
}
