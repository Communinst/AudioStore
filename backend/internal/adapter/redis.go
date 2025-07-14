package adapter

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	DB *redis.Client
}

func NewRedis(address,
	password string,
	db int) (*Redis, func(), error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db})

	err := conn.Ping(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("redis: connection failed: %w", err.Err())
	}

	return &Redis{DB: conn}, func() {
		if err := conn.Close(); err != nil {
			slog.Error("redis: failed to shutdown redis connection.", "Error:", err)
		}
	}, nil
}

func (r *Redis) Close() error {
	return r.DB.Close()
}
