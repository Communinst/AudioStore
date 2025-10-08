package minioAdapter

import (
	httpError "AudioShare/backend/internal/error"
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

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

func (m *MinioRepository) CreateBucket(ctx context.Context, bucketName string) error {
	exists, err := m.dbMinio.BucketExists(ctx, bucketName)
	if err != nil {
		return httpError.New(http.StatusInternalServerError,
			fmt.Sprintf("minio repository: create bucket: failed to obtain bucket info: %s", err.Error()))
	}

	if exists {
		slog.Info("minio repository: create bucket: %s bucket already exist.", bucketName)
		return nil
	}

	err = m.dbMinio.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
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
	reader := bytes.NewReader(data)
	_, err := m.dbMinio.PutObject(ctx, bucketName, objectName, reader, int64(len(data)), minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		slog.Error("minio repository: put object: failed to upload: %s", err.Error())
		return httpError.New(http.StatusInternalServerError,
			fmt.Sprintf("minio repository: put object: failed to upload: %s", err.Error()))
	}
	return nil
}

func (m *MinioRepository) GetObject(ctx context.Context, bucketName, objectName string) ([]byte, error) {
	object, err := m.dbMinio.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		slog.Error("minio repository: get object: failed to get: %s", err.Error())
		return nil, httpError.New(http.StatusInternalServerError,
			fmt.Sprintf("minio repository: get object: failed to get: %s", err.Error()))
	}
	defer object.Close()

	data, err := io.ReadAll(object)
	if err != nil {
		slog.Error("minio repository: get object: failed to read: %s", err.Error())
		return nil, httpError.New(http.StatusInternalServerError,
			fmt.Sprintf("minio repository: get object: failed to read: %s", err.Error()))
	}
	return data, nil
}

func (m *MinioRepository) GetObjectStream(ctx context.Context, bucketName, objectName string) (io.ReadCloser, int64, error) {
	object, err := m.dbMinio.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		slog.Error("minio repository: get object stream: failed to get: %s", err.Error())
		return nil, 0, httpError.New(http.StatusInternalServerError,
			fmt.Sprintf("minio repository: get object stream: failed to get: %s", err.Error()))
	}

	objInfo, err := m.dbMinio.StatObject(ctx, bucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		object.Close()
		slog.Error("minio repository: get object stream: failed to stat: %s", err.Error())
		return nil, 0, httpError.New(http.StatusInternalServerError,
			fmt.Sprintf("minio repository: get object stream: failed to stat: %s", err.Error()))
	}

	return object, objInfo.Size, nil
}

func (m *MinioRepository) RemoveObject(ctx context.Context, bucketName, objectName string) error {
	err := m.dbMinio.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		slog.Error("minio repository: remove object: failed to remove: %s", err.Error())
		return httpError.New(http.StatusInternalServerError,
			fmt.Sprintf("minio repository: remove object: failed to remove: %s", err.Error()))
	}
	return nil
}

func (m *MinioRepository) ObjectExists(ctx context.Context, bucketName, objectName string) (bool, error) {
	_, err := m.dbMinio.StatObject(ctx, bucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		if minio.ToErrorResponse(err).Code == "NoSuchKey" {
			return false, nil
		}
		slog.Error("minio repository: object exists: failed to check: %s", err.Error())
		return false, httpError.New(http.StatusInternalServerError,
			fmt.Sprintf("minio repository: object exists: failed to check: %s", err.Error()))
	}
	return true, nil
}

func (m *MinioRepository) PresignedGetObject(ctx context.Context, bucketName, objectName string, expirySec int) (string, error) {
	url, err := m.dbMinio.PresignedGetObject(ctx, bucketName, objectName, time.Duration(expirySec)*time.Second, nil)
	if err != nil {
		slog.Error("minio repository: presigned get object: failed to generate: %s", err.Error())
		return "", httpError.New(http.StatusInternalServerError,
			fmt.Sprintf("minio repository: presigned get object: failed to generate: %s", err.Error()))
	}
	return url.String(), nil
}

func (m *MinioRepository) PresignedPutObject(ctx context.Context, bucketName, objectName string, expirySec int) (string, error) {
	url, err := m.dbMinio.PresignedPutObject(ctx, bucketName, objectName, time.Duration(expirySec)*time.Second)
	if err != nil {
		slog.Error("minio repository: presigned put object: failed to generate: %s", err.Error())
		return "", httpError.New(http.StatusInternalServerError,
			fmt.Sprintf("minio repository: presigned put object: failed to generate: %s", err.Error()))
	}
	return url.String(), nil
}
