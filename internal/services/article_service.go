package services

import (
	"context"

	"portfolio-backend/internal/models"
	"portfolio-backend/internal/repositories"
)

type ArticleService struct {
	articleRepo *repositories.ArticleRepository
}

func NewArticleService(articleRepo *repositories.ArticleRepository) *ArticleService {
	return &ArticleService{articleRepo: articleRepo}
}

func (s *ArticleService) CreateArticle(ctx context.Context, article *models.Article) error {
	return s.articleRepo.CreateArticle(ctx, article)
}

func (s *ArticleService) GetArticles(ctx context.Context) ([]models.Article, error) {
	return s.articleRepo.GetArticles(ctx)
}