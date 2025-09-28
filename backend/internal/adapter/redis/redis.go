package redisAdapter

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	db *redis.Client
}

func NewRedis(host,
	port,
	password string,
	db int) (*RedisClient, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       db})

	err := conn.Ping(ctx).Err()
	if err != nil {
		return nil, fmt.Errorf("redis: connection failed: %w", err)
	}
	slog.Info("redis: connection established.")
	return &RedisClient{db: conn}, nil
}

func (r *RedisClient) Close() error {
	if err := r.db.Close(); err != nil {
		return fmt.Errorf("redis: failed to shutdown postgres connection: %w", err)
	}
	return nil
}
