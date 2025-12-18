package middleware

import (
	"go-rest-api/internal/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(jwtSecret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(utils.ErrorResponse("Missing authorization header"))
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(utils.ErrorResponse("Invalid authorization header format"))
		}

		token := parts[1]
		claims, err := utils.ValidateToken(token, jwtSecret)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(utils.ErrorResponse("Invalid or expired token"))
		}

		c.Locals("userID", claims.UserID)
		return c.Next()
	}
}