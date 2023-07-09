package auth

import (
	"context"

	"github.com/KadirbekSharau/Byte/internal/models"
)

type Repository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, email, password string) (*models.User, error)
}