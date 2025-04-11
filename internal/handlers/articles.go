package handlers

import (
	"github.com/gofiber/fiber/v2"

	"portfolio-backend/internal/models"
	"portfolio-backend/internal/services"
)

type ArticleHandler struct {
	articleService *services.ArticleService
}

func NewArticleHandler(articleService *services.ArticleService) *ArticleHandler {
	return &ArticleHandler{articleService: articleService}
}

func (h *ArticleHandler) CreateArticle(c *fiber.Ctx) error {
	var article models.Article
	if err := c.BodyParser(&article); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := h.articleService.CreateArticle(c.Context(), &article); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create article"})
	}

	return c.Status(fiber.StatusCreated).JSON(article)
}

func (h *ArticleHandler) GetArticles(c *fiber.Ctx) error {
	articles, err := h.articleService.GetArticles(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch articles"})
	}

	return c.JSON(articles)
}