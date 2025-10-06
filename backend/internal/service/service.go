package service

import (
	"AudioShare/backend/internal/entity"
	repository "AudioShare/backend/internal/repository/interfaces"
	"context"
)

const (
	auth_time_out = 15
)

type AuthServiceInteface interface {
	PostOne(ctx context.Context, data *entity.User) (int64, error)
	GetOneByEmail(ctx context.Context, email string) (*entity.User, error)
	GenerateAuthToken(user *entity.User, secret string, expireTime int) (string, error)
}

type EntityServiceInterface[E repository.Entity] interface {
	PostOne(ctx context.Context, data *E) (int64, error)
	GetOneById(ctx context.Context, id uint64) (*E, error)
	GetManyById(ctx context.Context, ids []uint64) ([]*E, error)
	GetAll(ctx context.Context) ([]*E, error)
	DeleteOneById(ctx context.Context, id uint64) error
	DeleteManyById(ctx context.Context, ids []uint64) (uint64, error)
}

type UserServiceInreface interface {
	EntityServiceInterface[entity.User]
}

type TrackServiceInterface interface {
	EntityServiceInterface[entity.TrackFile]
}

type Service struct {
	AuthServiceInteface
}

func NewService(
	postgres *repository.PostgresRepository,
	redis *repository.RedisRepository,
	minio *repository.MinioRepository) *Service {
	return &Service{
		//AuthServiceInteface: NewAuthService(postgres.AuthPostgresRepositoryInterface),
	}
}
