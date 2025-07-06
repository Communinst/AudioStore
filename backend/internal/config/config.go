package config

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Logger struct {
		Source   bool   `env:"LOGGER_SOURCE" env-default:"0"`
		Filepath string `env:"LOGGER_FILEPATH" env-default:"X:\\Coding\\AudioShare\\backend\\logs"`
	}
	HTTPServer struct {
		Address string        `env:"HTTP_ADDRESS" env-default:"localhost:8080"`
		Timeout time.Duration `env:"HTTP_TIMEOUT" env-default:"10s"`
	}
	SQL struct {
		Host     string `env:"SQL_HOST" env-default:"localhost"`
		Port     string `env:"SQL_PORT" env-default:"8080"`
		DBName   string `env:"SQL_DBNAME" env-default:"postgres"`
		Username string `env:"SQL_USERNAME" env-default:"admin"`
		Password string `env:"SQL_PASSWORD" env-required:"true"`
		SSLMode  bool   `env:"SQL_SSL" env-default:"0"`
	}
}

func (cfg *Config) Show() {
	fmt.Printf("Logger:\n\tSource: %t\n\tFilepath: %s", cfg.Logger.Source, cfg.Logger.Filepath)
}

func LoadConfig() *Config {
	var cfg Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		slog.Info("No .env was found.")
		return nil
	}
	cfg.Show()

	return nil``
}
