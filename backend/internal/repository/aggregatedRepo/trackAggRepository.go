package repositoryAggregated

import (
	"AudioShare/backend/internal/entity"
	repository "AudioShare/backend/internal/repository/interfaces"
	"context"
	"io"
)

type TrackAggregatedRepository struct {
	object repository.TrackMinioRepositoryInterface
}

func NewTrackAggregatedRepository(object repository.TrackMinioRepositoryInterface) *TrackAggregatedRepository {
	return &TrackAggregatedRepository{
		object: object,
	}
}

func (tar *TrackAggregatedRepository) CreateBucket(ctx context.Context, bucketName string) error {
	return tar.object.CreateBucket(ctx, bucketName)
}

func (tar *TrackAggregatedRepository) BucketExists(ctx context.Context, bucketName string) (bool, error) {
	return tar.object.BucketExists(ctx, bucketName)
}
func (tar *TrackAggregatedRepository) RemoveBucket(ctx context.Context, bucketName string) error {
	return tar.object.RemoveBucket(ctx, bucketName)
}

// // File operations
func (tar *TrackAggregatedRepository) PutObject(ctx context.Context, bucketName, objectName string, data []byte, contentType string) error {
	return tar.object.PutObject(ctx, bucketName, objectName, data, contentType)
}

func (tar *TrackAggregatedRepository) GetObject(ctx context.Context, bucketName, objectName string) ([]byte, error) {
	return tar.object.GetObject(ctx, bucketName, objectName)
}

func (tar *TrackAggregatedRepository) GetObjectStream(ctx context.Context, bucketName, objectName string) (io.ReadCloser, int64, error) {
	return tar.object.GetObjectStream(ctx, bucketName, objectName)
}
func (tar *TrackAggregatedRepository) RemoveObject(ctx context.Context, bucketName, objectName string) error {
	return tar.object.RemoveObject(ctx, bucketName, objectName)
}

func (tar *TrackAggregatedRepository) ObjectExists(ctx context.Context, bucketName, objectName string) (bool, error) {
	return tar.object.ObjectExists(ctx, bucketName, objectName)
}

// // Presigned URLs
func (tar *TrackAggregatedRepository) PresignedGetObject(ctx context.Context, bucketName, objectName string, expirySec int) (string, error) {
	return tar.object.PresignedGetObject(ctx, bucketName, objectName, expirySec)
}
func (tar *TrackAggregatedRepository) PresignedPutObject(ctx context.Context, bucketName, objectName string, expirySec int) (string, error) {
	return tar.object.PresignedGetObject(ctx, bucketName, objectName, expirySec)
}

func (tar *TrackAggregatedRepository) UploadTrack(ctx context.Context, req *entity.UploadRequest) (*entity.TrackFile, error) {
	return tar.object.UploadTrack(ctx, req)
}

func (tar *TrackAggregatedRepository) DownloadTrack(ctx context.Context, bucketName, objectName string) (*entity.DownloadResponse, error) {
	return tar.object.DownloadTrack(ctx, bucketName, objectName)
}

func (tar *TrackAggregatedRepository) StreamTrack(ctx context.Context, bucketName, objectName string, offset, length int64) (io.ReadCloser, int64, error) {
	return tar.object.StreamTrack(ctx, bucketName, objectName, offset, length)
}

func (tar *TrackAggregatedRepository) GetTrackInfo(ctx context.Context, bucketName, objectName string) (*entity.TrackFile, error) {
	return tar.object.GetTrackInfo(ctx, bucketName, objectName)
}

func (tar *TrackAggregatedRepository) ListTracks(ctx context.Context, bucketName, prefix string) ([]*entity.TrackFile, error) {
	return tar.object.ListTracks(ctx, bucketName, prefix)
}
func (tar *TrackAggregatedRepository) CopyTrack(ctx context.Context, srcBucket, srcObject, destBucket, destObject string) error {
	return tar.object.CopyTrack(ctx, srcBucket, srcBucket, destBucket, destObject)
}
