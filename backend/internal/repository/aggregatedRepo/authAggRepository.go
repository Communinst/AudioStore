package repositoryAggregated

import (
	"AudioShare/backend/internal/entity"
	repository "AudioShare/backend/internal/repository/interfaces"
	"context"
	"log/slog"
)

type AuthAggregatedRepository struct {
	db    repository.AuthPostgresRepositoryInterface
	cache repository.AuthRedisRepositoryInterface
}

func NewAggregatedRepository(pstgrs repository.AuthPostgresRepositoryInterface,
	rds repository.AuthRedisRepositoryInterface) *AuthAggregatedRepository {
	return &AuthAggregatedRepository{
		db:    pstgrs,
		cache: rds,
	}
}

func (this *AuthAggregatedRepository) PostOne(ctx context.Context, data *entity.User) (int64, error) {
	slog.Info("auth agg repository: post one: intitiated.")

	slog.Info("auth agg repository: post one: check cache.")
	result, err := this.cache.GetOneByEmail(ctx, data.Email)
	if err != nil {
		slog.Info(err.Error())
	} else if result != nil {
		return int64(result.Id), nil
	}

	slog.Info("auth agg repository: post one: request to db.")
	resultId, err := this.db.PostOne(ctx, data)
	if err == nil {
		return resultId, err
	}
	slog.Error(err.Error())
	return -1, err
}

// GetOneByEmail(ctx context.Context, email string) (*entity.User, error)
