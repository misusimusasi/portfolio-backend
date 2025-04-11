package handlers

import (
	"github.com/gofiber/fiber/v2"
	"portfolio-backend/internal/models"
	"portfolio-backend/internal/services"
	"portfolio-backend/pkg/jwt"
	"portfolio-backend/internal/config"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := h.authService.Register(c.Context(), &user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Registration failed"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created"})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	user, err := h.authService.Login(c.Context(), input.Username, input.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	cfg := config.Load()
	token, err := jwt.GenerateToken(user.ID, cfg.JWTSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate token"})
	}

	return c.JSON(fiber.Map{"token": token})
}