package repositoryAggregated

import (
	"AudioShare/backend/internal/entity"
	repository "AudioShare/backend/internal/repository/interfaces"
	"context"
	"io"
)

type AuthAggregatedRepositoryInterface interface {
	PostOne(ctx context.Context, data *entity.User) (int64, error)
	GetOneByEmail(ctx context.Context, email string) (*entity.UserCache, error)
	GetOneByEmailFull(ctx context.Context, email string) (*entity.User, error)
}

type DumpAggregatedRepositoryInterface interface {
	InsertDump(ctx context.Context, dump *entity.Dump) error
	GetAllDumps(ctx context.Context) ([]entity.Dump, error)
}

type EntityAggregatedRepositoryInterface[E repository.Entity] interface {
	PostOne(ctx context.Context, data *E) (int64, error)
	GetOneById(ctx context.Context, id uint64) (*E, error)
	GetAll(ctx context.Context) ([]*E, error)
	DeleteOneById(ctx context.Context, id uint64) error
}

type UserAggregatedRepositoryInterface interface {
	EntityAggregatedRepositoryInterface[entity.User]
	CheckIfUserWithRoleExists(ctx context.Context, roleId uint8) (bool, error)
}

type TrackAggregatedRepositoryInterface interface {
	CreateBucket(ctx context.Context, bucketName string) error
	BucketExists(ctx context.Context, bucketName string) (bool, error)
	RemoveBucket(ctx context.Context, bucketName string) error

	// File operations
	PutObject(ctx context.Context, bucketName, objectName string, data []byte, contentType string) error
	GetObject(ctx context.Context, bucketName, objectName string) ([]byte, error)
	GetObjectStream(ctx context.Context, bucketName, objectName string) (io.ReadCloser, int64, error)
	RemoveObject(ctx context.Context, bucketName, objectName string) error
	ObjectExists(ctx context.Context, bucketName, objectName string) (bool, error)

	// Presigned URLs
	PresignedGetObject(ctx context.Context, bucketName, objectName string, expirySec int) (string, error)
	PresignedPutObject(ctx context.Context, bucketName, objectName string, expirySec int) (string, error)

	UploadTrack(ctx context.Context, req *entity.UploadRequest) (*entity.TrackFile, error)
	DownloadTrack(ctx context.Context, bucketName, objectName string) (*entity.DownloadResponse, error)
	StreamTrack(ctx context.Context, bucketName, objectName string, offset, length int64) (io.ReadCloser, int64, error)
	GetTrackInfo(ctx context.Context, bucketName, objectName string) (*entity.TrackFile, error)

	ListTracks(ctx context.Context, bucketName, prefix string) ([]*entity.TrackFile, error)
	CopyTrack(ctx context.Context, srcBucket, srcObject, destBucket, destObject string) error
}

type AggregatedRepository struct {
	Auth  AuthAggregatedRepositoryInterface
	Dump  DumpAggregatedRepositoryInterface
	User  UserAggregatedRepositoryInterface
	Track TrackAggregatedRepositoryInterface
}

func NewAggregatedRepository(
	pstgrs repository.PostgresRepository,
	rds repository.RedisRepository,
	mn repository.MinioRepository) *AggregatedRepository {
	return &AggregatedRepository{
		Auth:  NewAuthAggregatedRepository(pstgrs.Auth, rds.Auth),
		Dump:  NewDumpAggregatedRepository(pstgrs.Dump),
		User:  NewUserAggregatedRepository(pstgrs.User, rds.User),
		Track: NewTrackAggregatedRepository(mn.Track),
	}
}
