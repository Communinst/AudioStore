package repository

import (
	"AudioShare/backend/internal/entity"
	"context"
	"time"
)

type AuthRedisRepositoryInterface interface {
}

type EntityRedisRepositoryInterface[E Entity] interface {
	BuildKey(key string) string
	PostOne(ctx context.Context, key string, data *entity.User, expiration time.Duration) error
	GetOne(ctx context.Context, key string) (*E, error)
	DeleteOne(ctx context.Context, key string) error
	DeletePattern(ctx context.Context, pattern string) error //Simplified clear option
}

type UserRedisRepositoryInterface interface {
	EntityRedisRepositoryInterface[entity.User]
	PostOneById(ctx context.Context, id uint64, data *entity.User, expiration time.Duration) error
	GetOneById(ctx context.Context, id uint64) (*entity.User, error)
	DeleteOneById(ctx context.Context, id uint64) error
}

type RedisRepository struct {
	UserRedisRepositoryInterface
}
