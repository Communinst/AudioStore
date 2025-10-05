package repositoryAggregated

import (
	"AudioShare/backend/internal/entity"
	repository "AudioShare/backend/internal/repository/interfaces"
	"context"
)

type AuthAggregatedRepositoryInterface interface {
	PostOne(ctx context.Context, data *entity.User) (int64, error)
	GetOneByEmail(ctx context.Context, email string) (*entity.User, error)
}

type EntityAggregatedRepositoryInterface[E repository.Entity] interface {
	PostOne(ctx context.Context, data *E) (int64, error)
	GetOneById(ctx context.Context, id uint64) (*E, error)
	GetAll(ctx context.Context) ([]*E, error)
	DeleteOneById(ctx context.Context, id uint64) error
}
