package repositoryAggregated

import (
	"AudioShare/backend/internal/entity"
	repository "AudioShare/backend/internal/repository/interfaces"
	"context"
)

type DumpAggregatedRepository struct {
	db repository.DumpPostgresRepositoryInterface
}

func NewDumpAggregatedRepository(db repository.DumpPostgresRepositoryInterface) *DumpAggregatedRepository {
	return &DumpAggregatedRepository{
		db: db,
	}
}

func (this *DumpAggregatedRepository) InsertDump(ctx context.Context, dump *entity.Dump) error {
	return this.db.InsertDump(ctx, dump)
}

func (this *DumpAggregatedRepository) GetAllDumps(ctx context.Context) ([]entity.Dump, error) {
	return this.db.GetAllDumps(ctx)
}
