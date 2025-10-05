package main

import (
	"AudioShare/backend/internal/adapter"
	postgresAdapter "AudioShare/backend/internal/adapter/postgres"
	redisAdapter "AudioShare/backend/internal/adapter/redis"
	"AudioShare/backend/internal/config"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Booting
	err := godotenv.Load()
	if err != nil {
		slog.Info("initial env file couldn't be reached.")
		return
	}
	cfg := config.LoadConfig(os.Getenv("CONN_CONFIG_PATH"))

	// Load migration here or right after repos inits?
	err = godotenv.Load(os.Getenv("MIGRATION_CONFIG_PATH"))
	if err != nil {
		slog.Info(",igration env file couldn't be reached.")
		return
	}

	postgreSQLConn := adapter.MustConnect(postgresAdapter.NewPostgres(cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
		cfg.Postgres.SSLMode))
	defer postgreSQLConn.Close() // not nil guaranteed

	redisConn := adapter.MustConnect(redisAdapter.NewRedis(cfg.Redis.Host,
		cfg.Redis.Port,
		cfg.Redis.Password,
		cfg.Redis.DBName)) // not nil guaranteed
	defer redisConn.Close()

}
