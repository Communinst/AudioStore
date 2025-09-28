package minioAdapter

import (
	httpError "AudioShare/backend/internal/error"
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/minio/minio-go/v7"
)

type MinioRepository struct {
	dbMinio *minio.Client
}

func NewMinioRepostitory(connWrapper *MinioClient) *MinioRepository {
	return &MinioRepository{
		dbMinio: connWrapper.db,
	}
}

func (m *MinioRepository) CreateBucket(ctx context.Context, bucketName, region string) error {
	exists, err := m.dbMinio.BucketExists(ctx, bucketName)
	if err != nil {
		return httpError.New(http.StatusInternalServerError,
			fmt.Sprintf("minio repository: create bucket: failed to obtain bucket info: %s", err.Error()))
	}

	if exists {
		slog.Info("minio repository: create bucket: %s bucket already exist.", bucketName)
		return nil
	}

	err = m.dbMinio.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: region})
	if err == nil {
		slog.Info("minio repository: create bucket: %s creation succeded.", bucketName)
		return nil
	}

	slog.Error("minio repository: create bucket: failed to create: %s", err.Error())
	return httpError.New(http.StatusInternalServerError,
		fmt.Sprintf("minio repository: create bucket: failed to create: %s", err.Error()))
}

func (m *MinioRepository) BucketExists(ctx context.Context, bucketName string) (bool, error) {
	exists, err := m.dbMinio.BucketExists(ctx, bucketName)
	if err != nil {
		slog.Error("minio repository: bucket exists: failed to obtain bucket info: %s", err.Error())
		return true, httpError.New(http.StatusInternalServerError,
			fmt.Sprintf("minio repository: bucket exists: failed to obtain bucket info: %s", err.Error()))
	}
	return exists, nil
}

func (m *MinioRepository) RemoveBucket(ctx context.Context, bucketName string) error {
	err := m.dbMinio.RemoveBucket(ctx, bucketName)
	if err == nil {
		return nil
	}
	slog.Error("minio repository: remove bucket: failed removal: %s", err.Error())
	return httpError.New(http.StatusInternalServerError,
		fmt.Sprintf("minio repository: remove bucket: failed removal: %s", err.Error()))
}

func (m *MinioRepository) PutObject(ctx context.Context, bucketName, objectName string, data []byte, contentType string) error {
	
	return nil
}

// GetObject(ctx context.Context, bucketName, objectName string) ([]byte, error)
// GetObjectStream(ctx context.Context, bucketName, objectName string) (io.ReadCloser, int64, error)
// RemoveObject(ctx context.Context, bucketName, objectName string) error
// ObjectExists(ctx context.Context, bucketName, objectName string) (bool, error)
