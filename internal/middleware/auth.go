package middleware

import (
	"strings"
	"github.com/gofiber/fiber/v2"
	"portfolio-backend/internal/config"
	"portfolio-backend/pkg/jwt"
)

func AuthMiddleware(c *fiber.Ctx) error {
    // 1. Получаем токен из заголовка
    authHeader := c.Get("Authorization")
    if authHeader == "" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Authorization header required",
        })
    }

    // 2. Извлекаем токен (формат: "Bearer <token>")
    tokenParts := strings.Split(authHeader, " ")
    if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid token format",
        })
    }

    // 3. Валидация токена
    cfg := config.Load()
    claims, err := jwt.ValidateToken(tokenParts[1], cfg.JWTSecret)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error":   "Invalid token",
            "details": err.Error(), // Добавляем детали для отладки
        })
    }

    // 4. Добавляем claims в контекст
    c.Locals("userID", claims.UserID)
    return c.Next()
}