package authRepository

import (
	"context"
	"errors"

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

func (*UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	return errors.New("")
}
func (*UserRepository) GetUser(ctx context.Context, email, password string) (*models.User, error) {
	return nil, errors.New("")
}