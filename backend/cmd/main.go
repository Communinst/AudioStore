package main

import (
	"AudioShare/backend/adapter"
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

	postgreSQL := adapter.NewPostgres(cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
		cfg.Postgres.SSLMode)

	postgreSQL.Db.DriverName()

}
