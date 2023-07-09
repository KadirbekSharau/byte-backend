package usecase

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/KadirbekSharau/Byte/internal/auth"
	"github.com/KadirbekSharau/Byte/internal/models"
	"github.com/dgrijalva/jwt-go/v4"
)

type AuthClaims struct {
	jwt.StandardClaims
	User *models.User `json:"user"`
}

type authUseCase struct {
	repo auth.Repository
	hashSalt string
	signingKey []byte
	expireDuration time.Duration
}

func NewAuthUseCase(repo auth.Repository, hashSalt string, signingKey []byte, expireDuration time.Duration) *authUseCase {
	return &authUseCase{
		repo: repo,
		hashSalt: hashSalt,
		signingKey: signingKey,
		expireDuration: expireDuration,
	}
}


func (a *authUseCase) SignUp(ctx context.Context, email, password string) error {
	pwd := sha256.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))

	user := &models.User{
		Email: email,
		Password: fmt.Sprintf("%x", pwd.Sum(nil)),
	}

	return a.repo.CreateUser(ctx, user)
}

func (a *authUseCase) SignIn(ctx context.Context, email, password string) (string, error) {
	pwd := sha256.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	user, err := a.repo.GetUser(ctx, email, password)
	if err != nil {
		return "", auth.ErrUserNotFound
	}

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.signingKey)
}

func (a *authUseCase) ParseToken(ctx context.Context, accessToken string) (*models.User, error) {
	token, err := jwt.ParseWithClaims(
		accessToken, 
		&AuthClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return a.signingKey, nil
		})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, auth.ErrInvalidAccessToken
}