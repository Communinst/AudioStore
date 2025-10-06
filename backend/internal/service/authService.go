package service

import (
	authToken "AudioShare/backend/internal/JSONWebTokens"
	"AudioShare/backend/internal/entity"
	repositoryAggregated "AudioShare/backend/internal/repository/aggregatedRepo"
	"context"
	"log/slog"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService struct {
	postgres repositoryAggregated.AuthAggregatedRepositoryInterface
}

func NewAuthService(p repositoryAggregated.AuthAggregatedRepositoryInterface) *AuthService {
	return &AuthService{
		postgres: p,
	}
}

func (this *AuthService) PostOne(ctx context.Context, data *entity.User) (int64, error) {
	slog.Info("auth service: sign up: initiated")

	ctx, cancel := context.WithTimeout(ctx, auth_time_out*time.Second)
	defer cancel()

	result, err := this.postgres.PostOne(ctx, data)
	slog.Info("auth service: sign up: finished")

	return result, err
}

func (this *AuthService) GetOneByEmail(ctx context.Context, email string) (*entity.User, error) {
	slog.Info("auth service: sign in: by email: initiated")

	ctx, cancel := context.WithTimeout(ctx, auth_time_out*time.Second)
	defer cancel()

	result, err := this.postgres.GetOneByEmail(ctx, email)
	slog.Info("auth service: sign in: by email: finished")

	return result, err
}

func (this *AuthService) GenerateAuthToken(user *entity.User, secret string, expireTime int) (string, error) {
	claims := &authToken.JWTToken{
		Email: user.Email,
		Id:    strconv.FormatUint(user.Id, 10),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expireTime))),
			Issuer:    "CWDB6Sem",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return result, nil
}
