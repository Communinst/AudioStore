package repository

import (
	"AudioShare/backend/internal/entity"
	"context"
)

type Entity interface {
	entity.User | entity.Role
}

type EntityRepository[E Entity] interface {
	PostOne(ctx context.Context, data *E) (int64, error)
	GetOneById(ctx context.Context, id uint64) (*E, error)
	GetManyById(ctx context.Context, ids []uint64) ([]*E, error)
	GetAll(ctx context.Context) ([]*E, error)
	DeleteOneById(ctx context.Context, id uint64) error
	DeleteManyById(ctx context.Context, ids []uint64) (uint64, error)
}

type UserRepositoryInterface interface {
	EntityRepository[entity.User]
}

type Repository interface {
	UserRepositoryInterface
}

