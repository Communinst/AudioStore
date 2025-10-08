package repository

import (
	minioAdapter "AudioShare/backend/internal/adapter/minio"
	"AudioShare/backend/internal/entity"
	"context"
	"io"
)

type MinioRepositoryInterface interface {
	// Bucket operations
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
}

type TrackMinioRepositoryInterface interface {
	MinioRepositoryInterface

	UploadTrack(ctx context.Context, req *entity.UploadRequest) (*entity.TrackFile, error)
	DownloadTrack(ctx context.Context, bucketName, objectName string) (*entity.DownloadResponse, error)
	StreamTrack(ctx context.Context, bucketName, objectName string, offset, length int64) (io.ReadCloser, int64, error)
	GetTrackInfo(ctx context.Context, bucketName, objectName string) (*entity.TrackFile, error)

	ListTracks(ctx context.Context, bucketName, prefix string) ([]*entity.TrackFile, error)
	CopyTrack(ctx context.Context, srcBucket, srcObject, destBucket, destObject string) error
}

type MinioRepository struct {
	Track TrackMinioRepositoryInterface
}

func NewMinioRepository(dbWrapper *minioAdapter.MinioClient) *MinioRepository {
	return &MinioRepository{
		Track: minioAdapter.NewTrackRepository(dbWrapper),
	}
}
