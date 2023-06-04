package authUsecase

import (
	"context"
	"errors"
	"time"

	"github.com/KadirbekSharau/Byte/internal/auth"
	"github.com/KadirbekSharau/Byte/internal/models"
)

type AuthUseCase struct {
	repo auth.UserRepository
	hashSalt string
	signingKey []byte
	expiredDuration time.Duration
}

func NewAuthUseCase(repo auth.UserRepository, hashSalt string, signingKey []byte, expiredDuration time.Duration) *AuthUseCase {
	return &AuthUseCase{
		repo: repo,
		hashSalt: hashSalt,
		signingKey: signingKey,
		expiredDuration: expiredDuration,
	}
}


func (*AuthUseCase) SignUp(ctx context.Context, username, password string) error {
	return errors.New("")
}

func (*AuthUseCase) SignIn(ctx context.Context, username, password string) (string, error) {
	return "", errors.New("")
}

func (*AuthUseCase) ParseToken(ctx context.Context, accessToken string) (*models.User, error) {
	return nil, errors.New("")
}