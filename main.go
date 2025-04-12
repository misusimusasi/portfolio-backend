package main

import (
	"context"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"portfolio-backend/internal/config"
	"portfolio-backend/internal/handlers"
	"portfolio-backend/internal/middleware"
	"portfolio-backend/internal/repositories"
	"portfolio-backend/internal/services"
	"portfolio-backend/pkg/database"
)

func main() {
	// Загрузка конфигурации
	cfg := config.Load()

	// Подключение к БД
	dbPool, err := database.NewPostgresPool(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer dbPool.Close()

	// Инициализация репозиториев
	userRepo := repositories.NewUserRepository(dbPool)
	articleRepo := repositories.NewArticleRepository(dbPool)

	// Инициализация сервисов
	authService := services.NewAuthService(userRepo)
	articleService := services.NewArticleService(articleRepo)

	// Инициализация обработчиков
	authHandler := handlers.NewAuthHandler(authService)
	articleHandler := handlers.NewArticleHandler(articleService)

	// Создание Fiber-приложения
	app := fiber.New()
	app.Use(cors.New())

	// Публичные роуты
	app.Post("/api/register", authHandler.Register)
	app.Post("/api/login", authHandler.Login)
	app.Get("/api/articles", articleHandler.GetArticles)

	// Защищённые роуты
	secured := app.Group("/api", middleware.AuthMiddleware)
	secured.Post("/articles", articleHandler.CreateArticle)

	// Запуск сервера
	log.Printf("Server starting on port %s", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
