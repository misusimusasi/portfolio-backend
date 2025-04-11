package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port       string
	JWTSecret  string
}

// Load загружает конфигурацию из .env файла и переменных окружения
func Load() *Config {
	// Загружаем .env файл (если существует)
	_ = godotenv.Load()

	cfg := &Config{
		DatabaseURL: getEnv("DATABASE_URL", ""),
		Port:       getEnv("PORT", "3001"),
		JWTSecret:  getEnv("JWT_SECRET", generateRandomSecret()),
	}

	// Валидация обязательных полей
	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is required in .env file or environment variables")
	}

	// Удаляем возможные кавычки в строке подключения
	cfg.DatabaseURL = strings.Trim(cfg.DatabaseURL, `"'`)

	return cfg
}

// getEnv получает переменную окружения с fallback значением
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
