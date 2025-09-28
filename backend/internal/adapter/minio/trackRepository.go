package minioAdapter

import "github.com/minio/minio-go/v7"

type TrackMinioRepository struct {
	db *minio.Client
}

func NewTrackRepository(dbWrapper *MinioClient) *TrackMinioRepository {
	return &TrackMinioRepository{
		db: dbWrapper.db,
	}
}

