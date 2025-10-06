package repositoryAggregated

import (
	"AudioShare/backend/internal/entity"
	repository "AudioShare/backend/internal/repository/interfaces"
	"context"
	"log/slog"
)

type AuthAggregatedRepository struct {
	db    repository.AuthPostgresRepositoryInterface
	cache repository.AuthRedisRepositoryInterface
}

func NewAuthAggregatedRepository(pstgrs repository.AuthPostgresRepositoryInterface,
	rds repository.AuthRedisRepositoryInterface) *AuthAggregatedRepository {
	return &AuthAggregatedRepository{
		db:    pstgrs,
		cache: rds,
	}
}

func (this *AuthAggregatedRepository) PostOne(ctx context.Context, data *entity.User) (int64, error) {
	slog.Info("auth agg repository: post one: intitiated.")

	slog.Info("auth agg repository: post one: check cache.")
	result, err := this.cache.GetOneByEmail(ctx, data.Email)
	if err != nil {
		slog.Info(err.Error())
	} else if result != nil {
		return int64(result.Id), nil
	}

	slog.Info("auth agg repository: post one: request to db.")
	resultId, err := this.db.PostOne(ctx, data)
	if err == nil {
		return resultId, err
	}
	slog.Error(err.Error())
	return -1, err
}

func (this *AuthAggregatedRepository) GetOneByEmail(ctx context.Context, email string) (*entity.UserCache, error) {
	slog.Info("auth agg repository: get one: by email: intitiated.")

	slog.Info("auth agg repository: get one: by email: check cache.")
	resultCache, err := this.cache.GetOneByEmail(ctx, email)
	if err != nil {
		slog.Info(err.Error())
	} else if resultCache != nil {
		return resultCache, nil
	}

	result, err := this.db.GetOneByEmail(ctx, email)
	if err == nil {
		if result != nil {
			slog.Info("auth agg repository: get one: by email: succeded.")
			return &entity.UserCache{
				Id:       result.Id,
				Email:    result.Email,
				Nickname: result.Nickname,
				RoleId:   result.RoleId,
			}, nil
		}
		slog.Info("auth agg repository: get one: by email: no data found.")
		return nil, nil
	}
	slog.Error("auth agg repository: get one: by email: failed.")
	return nil, err
}

func (this *AuthAggregatedRepository) GetOneByEmailFull(ctx context.Context, email string) (*entity.User, error) {
	slog.Info("auth agg repository: get one: by email: intitiated.")

	result, err := this.db.GetOneByEmail(ctx, email)
	if err == nil {
		if result != nil {
			slog.Info("auth agg repository: get one: by email: succeded.")
			return result, nil
		}
		slog.Info("auth agg repository: get one: by email: no data found.")
		return nil, nil
	}
	slog.Error("auth agg repository: get one: by email: failed.")
	return nil, err
}
