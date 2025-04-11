package repositories

// TODO
import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"

	"portfolio-backend/internal/models"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	_, err := r.db.Exec(
		ctx,
		"INSERT INTO users (username, password_hash) VALUES ($1, $2)",
		user.Username, user.PasswordHash,
	)
	return err
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow(
		ctx,
		"SELECT id, username, password_hash FROM users WHERE username = $1",
		username,
	).Scan(&user.ID, &user.Username, &user.PasswordHash)
	return &user, err
}