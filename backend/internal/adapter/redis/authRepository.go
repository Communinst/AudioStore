package redisAdapter

import (
	"AudioShare/backend/internal/entity"
	httpError "AudioShare/backend/internal/error"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type AuthRedisRepository struct {
	dbRedis *redis.Client
	prefix  string
}

func NewAuthRedisRepository(connWrapper *RedisClient) *AuthRedisRepository {
	return &AuthRedisRepository{
		dbRedis: connWrapper.db,
		prefix:  "user:",
	}
}

// //
func (u *AuthRedisRepository) BuildKey(key string) string {
	return u.prefix + key
}

func (u *AuthRedisRepository) Exists(ctx context.Context, key string) (bool, error) {
	fullKey := u.BuildKey(key)

	exists, err := u.dbRedis.Exists(ctx, fullKey).Result()
	if err != nil {
		slog.Error("user redis repository: exists: failed to check existence", slog.String("key", fullKey), slog.String("error", err.Error()))
		return false, httpError.New(http.StatusInternalServerError,
			fmt.Sprintf("user redis repository: exists: failed to check existence for key %s", fullKey))
	}

	return exists > 0, nil
}

////

// //
func (u *AuthRedisRepository) PostOne(ctx context.Context, key string, data *entity.User, expiration time.Duration) error {
	fullKey := u.BuildKey(key)

	jsonData, err := json.Marshal(data)
	if err != nil {
		slog.Error("user redis repository: post one: failed to marshal data", slog.String("error", err.Error()))
		return httpError.New(http.StatusInternalServerError,
			"user redis repository: post one: failed to marshal data")
	}

	err = u.dbRedis.Set(ctx, fullKey, jsonData, expiration).Err()
	if err != nil {
		slog.Error("user redis repository: post one: failed to post data", slog.String("key", fullKey), slog.String("error", err.Error()))
		return httpError.New(http.StatusInternalServerError,
			fmt.Sprintf("user redis repository: post one: failed to post data for key %s", fullKey))
	}

	slog.Info("user redis repository: post one: data posted successfully", slog.String("key", fullKey))
	return nil
}

func (u *AuthRedisRepository) PostOneById(ctx context.Context, data *entity.User, expiration time.Duration) error {
	slog.Info("user redis repository: post one by id: initiated.")
	return u.PostOne(ctx, fmt.Sprintf("id:%d", data.Id), data, expiration)
}

////

// // //
// func (u *AuthRedisRepository) GetOne(ctx context.Context, key string) (*entity.User, error) {
// 	fullKey := u.BuildKey(key)

// 	data, err := u.dbRedis.Get(ctx, fullKey).Bytes()
// 	if err != nil {
// 		if err == redis.Nil {
// 			slog.Info("user redis repository: get one: key not found", slog.String("key", fullKey))
// 			return nil, nil
// 		}
// 		slog.Error("user redis repository: get one: failed to get data", slog.String("key", fullKey), slog.String("error", err.Error()))
// 		return nil, httpError.New(http.StatusInternalServerError,
// 			fmt.Sprintf("user redis repository: get one: failed to get data for key %s", fullKey))
// 	}

// 	var user entity.User
// 	err = json.Unmarshal(data, &user)
// 	if err != nil {
// 		slog.Error("user redis repository: get one: failed to unmarshal data", slog.String("key", fullKey), slog.String("error", err.Error()))
// 		return nil, httpError.New(http.StatusInternalServerError,
// 			fmt.Sprintf("user redis repository: get one: failed to unmarshal data for key %s", fullKey))
// 	}

// 	slog.Info("user redis repository: get one: data retrieved successfully", slog.String("key", fullKey))
// 	return &user, nil
// }

// func (u *AuthRedisRepository) GetOneById(ctx context.Context, id uint64) (*entity.User, error) {
// 	slog.Info("user redis repository: get one by id: initiated.")
// 	return u.GetOne(ctx, fmt.Sprintf("id:%d", id))
// }

// ////

// // //
// func (u *AuthRedisRepository) DeleteOne(ctx context.Context, key string) error {
// 	fullKey := u.BuildKey(key)

// 	result, err := u.dbRedis.Del(ctx, fullKey).Result()
// 	if err != nil {
// 		slog.Error("user redis repository: delete one: failed to delete data", slog.String("key", fullKey), slog.String("error", err.Error()))
// 		return httpError.New(http.StatusInternalServerError,
// 			fmt.Sprintf("user redis repository: delete one: failed to delete data for key %s", fullKey))
// 	}

// 	if result > 0 {
// 		slog.Info("user redis repository: delete one: data deleted successfully", slog.String("key", fullKey))
// 	} else {
// 		slog.Info("user redis repository: delete one: key not found", slog.String("key", fullKey))
// 	}

// 	return nil
// }

// func (u *AuthRedisRepository) DeleteByID(ctx context.Context, id uint64) error {
// 	slog.Info("user redis repository: delete one by id: initiated.")
// 	return u.DeleteOne(ctx, fmt.Sprintf("%d", id))
// }

// func (u *AuthRedisRepository) DeletePattern(ctx context.Context, pattern string) error {

// 	keys, err := u.dbRedis.Keys(ctx, pattern).Result()
// 	if err != nil {
// 		slog.Error("user redis repository: delete pattern: failed to obtain pattern")
// 		return httpError.New(http.StatusInternalServerError,
// 			fmt.Sprintf("user redis repository: delete pattern: failed to obtain pattern"))
// 	}

// 	if len(keys) != 0 {
// 		_, err := u.dbRedis.Del(ctx, keys...).Result()
// 		if err != nil {
// 			slog.Error("user redis repository: delete pattern: failed to delete pattern")
// 			return httpError.New(http.StatusInternalServerError,
// 				fmt.Sprintf("user redis repository: delete pattern: failed to delete pattern"))
// 		}
// 	}

// 	return nil
// }
