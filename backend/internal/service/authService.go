package service

import (
	"AudioShare/backend/internal/entity"
	"AudioShare/backend/internal/repository/interfaces"
	"context"
	"log/slog"
	"time"
)

type AuthService struct {
	postgres repository.AuthPostgresRepositoryInterface
}

func NewAuthService(p repository.AuthPostgresRepositoryInterface) *AuthService {
	return &AuthService{
		postgres: p,
	}
}

func (this *AuthService) PostOne(ctx context.Context, data *entity.User) (int64, error) {
	slog.Info("auth service: post one: initiated")

	ctx, cancel := context.WithTimeout(ctx, auth_time_out*time.Second)
	defer cancel()

	result, err := this.postgres.PostOne(ctx, data)
	slog.Info("auth service: post one: finished")

	return result, err
}

func (this *AuthService) GetOneByEmail(ctx context.Context, email string) (*entity.User, error) {
	slog.Info("auth service: get one: by email: initiated")

	ctx, cancel := context.WithTimeout(ctx, auth_time_out*time.Second)
	defer cancel()

	result, err := this.postgres.GetOneByEmail(ctx, email)
	slog.Info("auth service: get one: by email: finished")

	return result, err
}
