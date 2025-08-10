package main

import (
	"AudioShare/backend/internal/adapter"
	"AudioShare/backend/internal/config"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// Booting
	err := godotenv.Load()
	if err != nil {
		slog.Info("initial env file couldn't be reached/")
		return
	}
	cfg := config.LoadConfig(os.Getenv("CONFIG_PATH"))

	postgreSQL := adapter.MustConnect(adapter.NewPostgres(cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
		cfg.Postgres.SSLMode))
	defer postgreSQL.Close() // not nil guaranteed

	redis := adapter.MustConnect(adapter.NewRedis(cfg.Redis.Host,
		cfg.Redis.Port,
		cfg.Redis.Password,
		cfg.Redis.DBName)) // not nil guaranteed
	defer redis.Close()


}
