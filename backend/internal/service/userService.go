package service

import (
	"AudioShare/backend/internal/entity"
	repositoryAggregated "AudioShare/backend/internal/repository/aggregatedRepo"
	"context"
)

type UserService struct {
	repo repositoryAggregated.UserAggregatedRepositoryInterface
}

func NewUserService(repo repositoryAggregated.UserAggregatedRepositoryInterface) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (this *UserService) PostOne(ctx context.Context, data *entity.User) (int64, error) {
	return this.repo.PostOne(ctx, data)
}
func (this *UserService) GetOneById(ctx context.Context, id uint64) (*entity.User, error) {
	return this.repo.GetOneById(ctx, id)
}
func (this *UserService) GetAll(ctx context.Context) ([]*entity.User, error) {
	return this.repo.GetAll(ctx)
}
func (this *UserService) DeleteOneById(ctx context.Context, id uint64) error {
	return this.repo.DeleteOneById(ctx, id)
}
func (this *UserService) CheckIfUserWithRoleExists(ctx context.Context, roleId uint8) (bool, error) {
	return this.repo.CheckIfUserWithRoleExists(ctx, roleId)
}
