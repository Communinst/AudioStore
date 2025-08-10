package adapter

import (
	"context"
	"fmt"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Minio struct {
	db *minio.Client
}

func NewMinio(host,
	port,
	region,
	endpoint,
	access,
	secret,
	bucket string,
	SSL bool) (*Minio, error) {

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

	return &Minio{db: conn}, nil
}
