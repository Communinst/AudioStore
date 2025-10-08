package service

import (
	"AudioShare/backend/internal/entity"
	repositoryAggregated "AudioShare/backend/internal/repository/aggregatedRepo"
	"context"
)

type trackService struct {
	repo repositoryAggregated.TrackAggregatedRepositoryInterface
}

func NewTrackService(repo repositoryAggregated.TrackAggregatedRepositoryInterface) TrackServiceInterface {
	return &trackService{
		repo: repo,
	}
}

func (t *trackService) UploadTrack(ctx context.Context, req *entity.UploadRequest) (uint64, error) {
	// TODO: Implement track upload logic
	return 0, nil
}

func (t *trackService) DownloadTrack(ctx context.Context, bucket string, objectKey string) (entity.DownloadResponse, error) {
	// TODO: Implement track download logic
	return entity.DownloadResponse{}, nil
}

func (t *trackService) GetTrackInfo(ctx context.Context, bucket string, objectKey string) (entity.TrackFile, error) {
	// TODO: Implement get track info logic
	return entity.TrackFile{}, nil
}
