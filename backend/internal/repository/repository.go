package repository

import (
	"AudioShare/backend/internal/entity"
	"context"
)

type Entity interface {
	entity.User | entity.Role
}

type EntityRepository[E Entity] interface {
	PostOne(ctx context.Context, data *E) (int, error)
	GetOneById(ctx context.Context, id int) (*E, error)
	GetManyById(ctx context.Context, id []int) ([]*E, error)
	DeleteOneById(ctx context.Context, id int) error
	DeleteManyById(ctx context.Context, id []int) error
}

type UserRepositoryInterface interface {
	EntityRepository[entity.User]
	UpdateRoleById(ctx context.Context, userId int, roleId int) (int, error)
}

type Repository interface {
	UserRepository
}
