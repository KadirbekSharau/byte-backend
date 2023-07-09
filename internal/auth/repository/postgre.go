package repository

import (
	"context"

	"github.com/KadirbekSharau/Byte/internal/models"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (email, password) VALUES (:email, :password)`
	_, err := r.db.NamedExecContext(ctx, query, user)
	return err
}

func (r *userRepository) GetUser(ctx context.Context, email, password string) (*models.User, error) {
	query := "SELECT * FROM users WHERE email = $1"
	user := &models.User{}
	err := r.db.GetContext(ctx, user, query, email)

	return user, err
}
