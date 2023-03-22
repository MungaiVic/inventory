package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func Protected() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte("secret"),
		ErrorHandler: jwtError,
	})
}

func jwtError(ctx *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Missing or malformed JWT",
		})
	} else {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": "error", 
			"message": "Invalid or expired JWT",
		})
	}
}
