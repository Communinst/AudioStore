package repositoryAggregated

import (
	"AudioShare/backend/internal/entity"
	repository "AudioShare/backend/internal/repository/interfaces"
	"context"
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

func PostOne(ctx context.Context, data *entity.User) (int64, error) {
	return 0, nil
}

// GetOneByEmail(ctx context.Context, email string) (*entity.User, error)
