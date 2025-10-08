package minioAdapter

import (
	"context"
	"fmt"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	db *minio.Client
}

func NewMinio(host,
	port,
	region,
	endpoint,
	access,
	secret,
	bucket string,
	SSL bool) (*MinioClient, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(access, secret, ""),
		Secure: SSL,
	})

	if err != nil {
		return nil, fmt.Errorf("minio: connection establishment failed: %w", err)
	}

	err = conn.MakeBucket(ctx, bucket, minio.MakeBucketOptions{Region: region})
	if err != nil {
		return nil, fmt.Errorf("minio: bucket creation failed: %w", err)
	}

	return &MinioClient{db: conn}, nil
}

func (m *MinioClient) Close() error {
	return nil
}
