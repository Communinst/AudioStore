package main

import (
	"AudioShare/backend/internal/adapter"
	"AudioShare/backend/internal/config"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func setSlog()

func main() {

	err := godotenv.Load()
	if err != nil {
		slog.Info("initial env file couldn't be reached/")
		return
	}

	cfg := config.LoadConfig(os.Getenv("CONFIG_PATH"))
	cfg.Show()

	postgreSQL, postgresCleanUp, err := adapter.NewPostgres(cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
		cfg.Postgres.SSLMode)
	redis, redisCleanUp, err := adapter.NewRedis(cfg.Redis.Host,
		cfg.Redis.Port,
		cfg.Redis.Password,
		cfg.Redis.DBName)

	postgreSQL.DB.DriverName()

	defer postgresCleanUp()
	defer redisCleanUp()

}
