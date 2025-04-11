package middleware

import (
	"github.com/gofiber/fiber/v2"
	"portfolio-backend/internal/config"
	"portfolio-backend/pkg/jwt"
)

func AuthMiddleware(c *fiber.Ctx) error {
	cfg := config.Load()
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	claims, err := jwt.ValidateToken(token, cfg.JWTSecret) // Добавляем второй аргумент
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	c.Locals("userID", claims.UserID)
	return c.Next()
}