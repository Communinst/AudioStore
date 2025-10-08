package minioAdapter

import (
	"AudioShare/backend/internal/entity"
	"bytes"
	"context"
	"fmt"
	"io"
	"path/filepath"
	"strconv"
	"time"

	"github.com/minio/minio-go/v7"
)

type TrackMinioRepository struct {
	db *minio.Client
}

func NewTrackRepository(dbWrapper *MinioClient) *TrackMinioRepository {
	return &TrackMinioRepository{
		db: dbWrapper.db,
	}
}

// Track-specific methods
func (r *TrackMinioRepository) UploadTrack(ctx context.Context, req *entity.UploadRequest) (*entity.TrackFile, error) {
	// Generate unique object key
	objectKey := fmt.Sprintf("tracks/%d/%s", req.UserID, req.FileName)

	// Upload to MinIO
	reader := bytes.NewReader(req.FileData)
	_, err := r.db.PutObject(ctx, "mybucket", objectKey, reader, int64(len(req.FileData)), minio.PutObjectOptions{
		ContentType: req.ContentType,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to upload track: %w", err)
	}

	// Create TrackFile entity
	track := &entity.TrackFile{
		ID:          0, // Will be set by database if needed
		Bucket:      "mybucket",
		ObjectKey:   objectKey,
		FileName:    req.FileName,
		ContentType: req.ContentType,
		Size:        int64(len(req.FileData)),
		UploadedAt:  time.Now(),
		UploaderId:  req.UserID,
		TrackName:   req.OriginalName,
	}

	return track, nil
}

func (r *TrackMinioRepository) DownloadTrack(ctx context.Context, bucketName, objectName string) (*entity.DownloadResponse, error) {
	object, err := r.db.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get object: %w", err)
	}
	defer object.Close()

	data, err := io.ReadAll(object)
	if err != nil {
		return nil, fmt.Errorf("failed to read object: %w", err)
	}

	// Get object info for metadata
	objInfo, err := r.db.StatObject(ctx, bucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get object info: %w", err)
	}

	return &entity.DownloadResponse{
		FileData:    data,
		FileName:    filepath.Base(objectName),
		ContentType: objInfo.ContentType,
		Size:        objInfo.Size,
	}, nil
}

func (r *TrackMinioRepository) StreamTrack(ctx context.Context, bucketName, objectName string, offset, length int64) (io.ReadCloser, int64, error) {
	object, err := r.db.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get object stream: %w", err)
	}

	objInfo, err := r.db.StatObject(ctx, bucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		object.Close()
		return nil, 0, fmt.Errorf("failed to get object info: %w", err)
	}

	return object, objInfo.Size, nil
}

func (r *TrackMinioRepository) GetTrackInfo(ctx context.Context, bucketName, objectName string) (*entity.TrackFile, error) {
	objInfo, err := r.db.StatObject(ctx, bucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get object info: %w", err)
	}

	// Extract user ID from object key (assuming format: tracks/{userID}/{filename})
	userID := uint64(0)
	if len(objInfo.Key) > 7 && objInfo.Key[:7] == "tracks/" {
		parts := objInfo.Key[7:]
		for idx := 0; idx < len(parts); idx++ {
			if parts[idx] == '/' {
				if id, err := strconv.ParseUint(parts[:idx], 10, 64); err == nil {
					userID = id
				}
				break
			}
		}
	}

	return &entity.TrackFile{
		ID:          0, // Would need database lookup for real ID
		Bucket:      bucketName,
		ObjectKey:   objectName,
		FileName:    filepath.Base(objectName),
		ContentType: objInfo.ContentType,
		Size:        objInfo.Size,
		UploadedAt:  objInfo.LastModified,
		UploaderId:  userID,
		TrackName:   filepath.Base(objectName),
	}, nil
}

func (r *TrackMinioRepository) ListTracks(ctx context.Context, bucketName, prefix string) ([]*entity.TrackFile, error) {
	objectCh := r.db.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
		Prefix:    prefix,
		Recursive: true,
	})

	var tracks []*entity.TrackFile
	for object := range objectCh {
		if object.Err != nil {
			return nil, fmt.Errorf("failed to list objects: %w", object.Err)
		}

		// Extract user ID from object key
		userID := uint64(0)
		if len(object.Key) > 7 && object.Key[:7] == "tracks/" {
			parts := object.Key[7:]
			for idx := 0; idx < len(parts); idx++ {
				if parts[idx] == '/' {
					if id, err := strconv.ParseUint(parts[:idx], 10, 64); err == nil {
						userID = id
					}
					break
				}
			}
		}

		track := &entity.TrackFile{
			ID:          0, // Would need database lookup for real ID
			Bucket:      bucketName,
			ObjectKey:   object.Key,
			FileName:    filepath.Base(object.Key),
			ContentType: "audio/mpeg", // Default, would need to get from object metadata
			Size:        object.Size,
			UploadedAt:  object.LastModified,
			UploaderId:  userID,
			TrackName:   filepath.Base(object.Key),
		}
		tracks = append(tracks, track)
	}

	return tracks, nil
}

func (r *TrackMinioRepository) CopyTrack(ctx context.Context, srcBucket, srcObject, destBucket, destObject string) error {
	_, err := r.db.CopyObject(ctx, minio.CopyDestOptions{
		Bucket: destBucket,
		Object: destObject,
	}, minio.CopySrcOptions{
		Bucket: srcBucket,
		Object: srcObject,
	})
	return err
}

// Implement MinioRepositoryInterface methods
func (r *TrackMinioRepository) CreateBucket(ctx context.Context, bucketName string) error {
	exists, err := r.db.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	return r.db.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
}

func (r *TrackMinioRepository) BucketExists(ctx context.Context, bucketName string) (bool, error) {
	return r.db.BucketExists(ctx, bucketName)
}

func (r *TrackMinioRepository) RemoveBucket(ctx context.Context, bucketName string) error {
	return r.db.RemoveBucket(ctx, bucketName)
}

func (r *TrackMinioRepository) PutObject(ctx context.Context, bucketName, objectName string, data []byte, contentType string) error {
	reader := bytes.NewReader(data)
	_, err := r.db.PutObject(ctx, bucketName, objectName, reader, int64(len(data)), minio.PutObjectOptions{
		ContentType: contentType,
	})
	return err
}

func (r *TrackMinioRepository) GetObject(ctx context.Context, bucketName, objectName string) ([]byte, error) {
	object, err := r.db.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer object.Close()
	return io.ReadAll(object)
}

func (r *TrackMinioRepository) GetObjectStream(ctx context.Context, bucketName, objectName string) (io.ReadCloser, int64, error) {
	object, err := r.db.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, 0, err
	}

	objInfo, err := r.db.StatObject(ctx, bucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		object.Close()
		return nil, 0, err
	}

	return object, objInfo.Size, nil
}

func (r *TrackMinioRepository) RemoveObject(ctx context.Context, bucketName, objectName string) error {
	return r.db.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
}

func (r *TrackMinioRepository) ObjectExists(ctx context.Context, bucketName, objectName string) (bool, error) {
	_, err := r.db.StatObject(ctx, bucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		if minio.ToErrorResponse(err).Code == "NoSuchKey" {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *TrackMinioRepository) PresignedGetObject(ctx context.Context, bucketName, objectName string, expirySec int) (string, error) {
	presignedURL, err := r.db.PresignedGetObject(ctx, bucketName, objectName, time.Duration(expirySec)*time.Second, nil)
	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}

func (r *TrackMinioRepository) PresignedPutObject(ctx context.Context, bucketName, objectName string, expirySec int) (string, error) {
	presignedURL, err := r.db.PresignedPutObject(ctx, bucketName, objectName, time.Duration(expirySec)*time.Second)
	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}
