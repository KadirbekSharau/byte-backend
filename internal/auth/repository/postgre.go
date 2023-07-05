package authRepository

import (
	"context"

	"github.com/KadirbekSharau/Byte/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (email, password) VALUES (:email, :password)`
	_, err := r.db.NamedExecContext(ctx, query, user)
	return err
}

func (r *UserRepository) GetUser(ctx context.Context, email, password string) (*models.User, error) {
	query := "SELECT * FROM users WHERE email = $1"
	user := &models.User{}
	err := r.db.GetContext(ctx, user, query, email)

	return user, err
}
