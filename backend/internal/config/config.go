package config

import (
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
	HTTPServer struct {
		Address string        `env:"HTTP_ADDRESS" env-default:"localhost:8080"`
		Timeout time.Duration `env:"HTTP_TIMEOUT" env-default:"10s"`
	}
	Postgres struct {
		Host     string `env:"POSTGRES_HOST" env-default:"localhost"`
		Port     string `env:"POSTGRES_PORT" env-default:"8080"`
		DBName   string `env:"POSTGRES_DBNAME" env-default:"postgres"`
		Username string `env:"POSTGRES_USER" env-default:"admin"`
		Password string `env:"POSTGRES_PASSWORD" env-required:"true"`
		SSLMode  string `env:"POSTGRES_SSL" env-default:"disable"`
	}
	Redis struct {
		Host     string `env:"REDIS_HOST" env-default:"localhost"`
		Port     string `env:"REDIS_PORT" env-default:"6379"`
		Password string `env:"REDIS_PASSWORD" env-required:"true"`
		DBName   int    `env:"REDIS_DB" env-default:"0"`
	}
	Minio struct {
		Host       string `env:"MINIO_HOST" env-default:"localhost"`
		Port       string `env:"MINIO_PORT" env-default:"9000"`
		Region     string `env:"MINIO_REGION" env-default:"us-east-1"`
		Endpoint   string `env:"MINIO_ENDPOINT" env-default:"http://localhost:9000"`
		AccessKey  string `env:"MINIO_ACCESS" env-required:"true"`
		SecretKey  string `env:"MINIO_SECRET" env-required:"true"`
		BucketName string `env:"MINIO_BUCKETNAME" env-required:"true"`
		SSLMode    bool   `env:"MINIO_SSL" env-default:"0"`
	}
}

// func (cfg *Config) Show() {
// 	fmt.Printf("%+v\n", *cfg)
// }

func LoadConfig(cfg_name string) *Config {
	var cfg Config
	err := cleanenv.ReadConfig(cfg_name, &cfg)
	if err != nil {
		slog.Info("Wrong config path.")
		return nil
	}
	return &cfg
}
