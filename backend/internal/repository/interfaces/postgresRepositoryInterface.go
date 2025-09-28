package repository

import (
	postgresAdapter "AudioShare/backend/internal/adapter/postgres"
	"AudioShare/backend/internal/entity"
	"context"
)

type EntityPostgresRepository[E Entity] interface {
	PostOne(ctx context.Context, data *E) (int64, error)
	GetOneById(ctx context.Context, id uint64) (*E, error)
	GetAll(ctx context.Context) ([]*E, error)
	DeleteOneById(ctx context.Context, id uint64) error
}

type AuthPostgresRepositoryInterface interface {
	PostOne(ctx context.Context, data *entity.User) (int64, error)
	GetOneByEmail(ctx context.Context, email string) (*entity.User, error)
}

type UserPostgresRepositoryInterface interface {
	EntityPostgresRepository[entity.User]
}

type PostgresRepository struct {
	auth AuthPostgresRepositoryInterface
	user UserPostgresRepositoryInterface
}

func NewPostgresRepository(dbWrapper *postgresAdapter.PostgresClient) *PostgresRepository {
	return &PostgresRepository{
		auth: postgresAdapter.NewAuthPostgresRepository(dbWrapper),
		user: postgresAdapter.NewUserPostgresRepository(dbWrapper),
	}
}
