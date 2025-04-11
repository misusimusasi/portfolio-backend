package repositories

// TODO
import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"

	"portfolio-backend/internal/models"
)

type ArticleRepository struct {
	db *pgxpool.Pool
}

func NewArticleRepository(db *pgxpool.Pool) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (r *ArticleRepository) CreateArticle(ctx context.Context, article *models.Article) error {
	_, err := r.db.Exec(
		ctx,
		"INSERT INTO articles (title, content, image_url) VALUES ($1, $2, $3)",
		article.Title, article.Content, article.ImageURL,
	)
	return err
}

func (r *ArticleRepository) GetArticles(ctx context.Context) ([]models.Article, error) {
	rows, err := r.db.Query(ctx, "SELECT * FROM articles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var a models.Article
		err := rows.Scan(&a.ID, &a.Title, &a.Content, &a.ImageURL, &a.CreatedAt)
		if err != nil {
			return nil, err
		}
		articles = append(articles, a)
	}
	return articles, nil
}