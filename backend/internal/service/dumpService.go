package service

import (
	"AudioShare/backend/internal/entity"
	repositoryAggregated "AudioShare/backend/internal/repository/aggregatedRepo"
	"context"
)

type DumpService struct {
	repo repositoryAggregated.DumpAggregatedRepositoryInterface
}

func NewDumpService(posgres repositoryAggregated.DumpAggregatedRepositoryInterface) *DumpService {
	return &DumpService{repo: posgres}
}

func (s *DumpService) InsertDump(ctx context.Context, filePath string, size int64) error {
	dump := &entity.Dump{
		Filename: filePath,
		Size:     size,
	}
	return s.repo.InsertDump(ctx, dump)
}

func (s *DumpService) GetAllDumps(ctx context.Context) ([]entity.Dump, error) {
	return s.repo.GetAllDumps(ctx)
}
