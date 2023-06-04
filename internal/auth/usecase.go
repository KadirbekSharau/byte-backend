package auth

import (
	"context"

	"github.com/KadirbekSharau/Byte/internal/models"
)

type UseCase interface {
	SignUp(ctx context.Context, username, password string) error
	SignIn(ctx context.Context, username, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*models.User, error)
}