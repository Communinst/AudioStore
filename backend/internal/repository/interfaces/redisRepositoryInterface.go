package repository

import (
	"AudioShare/backend/internal/entity"
	"context"
	"time"
)

type AuthRedisRepositoryInterface interface {
	PostOneById(ctx context.Context, data *entity.UserCache, expiration time.Duration) error
	PostOneByEmail(ctx context.Context, data *entity.UserCache, expiration time.Duration) error
	GetOneById(ctx context.Context, id uint64) (*entity.UserCache, error)
	GetOneByEmail(ctx context.Context, email string) (*entity.UserCache, error)
	DeleteByID(ctx context.Context, id uint64) error
	DeleteByEmail(ctx context.Context, email string) error
	DeletePattern(ctx context.Context, pattern string) error
}

type EntityRedisRepositoryInterface[E Entity] interface {
	BuildKey(key string) string
	PostOne(ctx context.Context, key string, data *E, expiration time.Duration) error
	GetOne(ctx context.Context, key string) (*E, error)
	DeleteOne(ctx context.Context, key string) error
	DeletePattern(ctx context.Context, pattern string) error //Simplified clear option
}

type UserRedisRepositoryInterface interface {
	EntityRedisRepositoryInterface[entity.UserCache]
	PostOneById(ctx context.Context, id uint64, data *entity.UserCache, expiration time.Duration) error
	GetOneById(ctx context.Context, id uint64) (*entity.UserCache, error)
	DeleteOneById(ctx context.Context, id uint64) error
}

type RedisRepository struct {
	UserRedisRepositoryInterface
}
