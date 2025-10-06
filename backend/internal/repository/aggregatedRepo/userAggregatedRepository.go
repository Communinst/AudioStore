package repositoryAggregated

import (
	"AudioShare/backend/internal/entity"
	repository "AudioShare/backend/internal/repository/interfaces"
	"context"
)

type UserAggregatedRepository struct {
	db    repository.UserPostgresRepositoryInterface
	cache repository.UserRedisRepositoryInterface
}

func NewUserAggregatedRepository(
	db repository.UserPostgresRepositoryInterface,
	cache repository.UserRedisRepositoryInterface) *UserAggregatedRepository {
	return &UserAggregatedRepository{
		db:    db,
		cache: cache,
	}
}

func (this *UserAggregatedRepository) PostOne(ctx context.Context, data *entity.User) (int64, error) {
	return this.db.PostOne(ctx, data)
}

func (this *UserAggregatedRepository) GetOneById(ctx context.Context, id uint64) (*entity.User, error) {
	return this.db.GetOneById(ctx, id)
}

func (this *UserAggregatedRepository) GetAll(ctx context.Context) ([]*entity.User, error) {
	return this.db.GetAll(ctx)
}

func (this *UserAggregatedRepository) DeleteOneById(ctx context.Context, id uint64) error {
	return this.db.DeleteOneById(ctx, id)
}

func (this *UserAggregatedRepository) CheckIfUserWithRoleExists(ctx context.Context, roleId uint8) (bool, error) {
	return this.db.CheckIfUserWithRoleExists(ctx, roleId)
}
