package repositoryAggregated

import (
	"AudioShare/backend/internal/entity"
	repository "AudioShare/backend/internal/repository/interfaces"
	"context"
)

type AuthAggregatedRepositoryInterface interface {
	PostOne(ctx context.Context, data *entity.User) (int64, error)
	GetOneByEmail(ctx context.Context, email string) (*entity.UserCache, error)
}

type DumpAggregatedRepositoryInterface interface {
	InsertDump(ctx context.Context, dump *entity.Dump) error
	GetAllDumps(ctx context.Context) ([]entity.Dump, error)
}

type EntityAggregatedRepositoryInterface[E repository.Entity] interface {
	PostOne(ctx context.Context, data *E) (int64, error)
	GetOneById(ctx context.Context, id uint64) (*E, error)
	GetAll(ctx context.Context) ([]*E, error)
	DeleteOneById(ctx context.Context, id uint64) error
}

type UserAggregatedRepositoryInterface interface {
	EntityAggregatedRepositoryInterface[entity.User]
	CheckIfUserWithRoleExists(ctx context.Context, roleId uint8) (bool, error)
}

type AggregatedRepository struct {
	Auth AuthAggregatedRepositoryInterface
	Dump DumpAggregatedRepositoryInterface
	User UserAggregatedRepositoryInterface
}

func NewAggregatedRepository(
	pstgrs repository.PostgresRepository,
	rds repository.RedisRepository,
	mn repository.MinioRepository) *AggregatedRepository {
	return &AggregatedRepository{
		Auth: NewAuthAggregatedRepository(pstgrs.Auth, rds.Auth),
		Dump: NewDumpAggregatedRepository(pstgrs.Dump),
		User: NewUserAggregatedRepository(pstgrs.User, rds.User),
	}
}
