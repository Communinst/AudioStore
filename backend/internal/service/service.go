package service

import (
	"AudioShare/backend/internal/entity"
	repositoryAggregated "AudioShare/backend/internal/repository/aggregatedRepo"
	repository "AudioShare/backend/internal/repository/interfaces"
	"context"
)

const (
	auth_time_out = 15
)

type AuthServiceInterface interface {
	PostOne(ctx context.Context, data *entity.User) (int64, error)
	GetOneByEmail(ctx context.Context, email string) (*entity.UserCache, error)
	GenerateAuthToken(user *entity.User, secret string, expireTime int) (string, error)
	GetOneByEmailFull(ctx context.Context, email string) (*entity.User, error)
}

type DumpServiceInterface interface {
	InsertDump(ctx context.Context, filePath string, size int64) error
	GetAllDumps(ctx context.Context) ([]entity.Dump, error)
}

type EntityServiceInterface[E repository.Entity] interface {
	PostOne(ctx context.Context, data *E) (int64, error)
	GetOneById(ctx context.Context, id uint64) (*E, error)
	GetAll(ctx context.Context) ([]*E, error)
	DeleteOneById(ctx context.Context, id uint64) error
}

type UserServiceInterface interface {
	EntityServiceInterface[entity.User]
	CheckIfUserWithRoleExists(ctx context.Context, roleId uint8) (bool, error)
}

type TrackServiceInterface interface { // define return types
	UploadTrack(ctx context.Context, req *entity.UploadRequest) (uint64, error)
	DownloadTrack(ctx context.Context, bucket string, objectKey string) (entity.DownloadResponse, error)
	GetTrackInfo(ctx context.Context, bucket string, objectKey string) (entity.TrackFile, error)
}

type Service struct {
	Auth  AuthServiceInterface
	Dump  DumpServiceInterface
	User  UserServiceInterface
	Track TrackServiceInterface
}

func NewService(repo *repositoryAggregated.AggregatedRepository) *Service {
	return &Service{
		Auth:  NewAuthService(repo.Auth),
		Dump:  NewDumpService(repo.Dump),
		User:  NewUserService(repo.User),
		Track: NewTrackService(),
	}
}
