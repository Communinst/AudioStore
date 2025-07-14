package config

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type EnvPath struct {
	Config_path string `env:"CONFIG_PATH" env-required:"true"`
}

func LoadEnv() *EnvPath {
	var paths EnvPath
	err := cleanenv.ReadEnv(&paths)
	if err != nil {
		slog.Info("no .env's found.")
		return nil
	}

	return &paths
}

type Config struct {
	// Logger struct {
	// 	Source bool `env:"LOGGER_SOURCE" env-default:"0"`
	// 	//Filepath string `env:"LOGGER_FILEPATH" env-default:"./logs"`
	// 	Level int `env:"LOGGER_LEVEL" env-default:"-4"`
	// }
	HTTPServer struct {
		Address string        `env:"HTTP_ADDRESS" env-default:"localhost:8080"`
		Timeout time.Duration `env:"HTTP_TIMEOUT" env-default:"10s"`
	}
	Postgres struct {
		Host     string `env:"POSTGRES_HOST" env-default:"localhost"`
		Port     string `env:"POSTGRES_PORT" env-default:"8080"`
		DBName   string `env:"POSTGRES_DBNAME" env-default:"postgres"`
		Username string `env:"POSTGRES_USERNAME" env-default:"admin"`
		Password string `env:"POSTGRES_PASSWORD" env-required:"true"`
		SSLMode  string `env:"POSTGRES_SSL" env-default:"disable"`
	}
	Redis struct {
		Host     string `env:"REDIS_HOST" env-default:"localhost"`
		Port     string `env:"REDIS_PORT" env-default:"6379"`
		Password string `env:"REDIS_PASSWORD" env-required:"true"`
		DBName   int    `env:"REDIS_DB" env-default:"0"`
	}
}

func (cfg *Config) Show() {
	fmt.Printf("%+v\n", *cfg)
}

func LoadConfig(cfg_name string) *Config {
	var cfg Config
	err := cleanenv.ReadConfig(cfg_name, &cfg)
	if err != nil {
		slog.Info("Wrong config path.")
		return nil
	}
	return &cfg
}
