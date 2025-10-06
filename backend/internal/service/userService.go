package service

import (
	repositoryAggregated "AudioShare/backend/internal/repository/aggregatedRepo"
)

type UserService struct {
	repo repositoryAggregated.UserAggregatedRepositoryInterface
}

func NewUserService(repo repositoryAggregated.UserAggregatedRepositoryInterface) *UserService {
	return &UserService{
		repo: repo,
	}
}
