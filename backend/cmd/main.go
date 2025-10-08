package main

import (
	"AudioShare/backend/internal/adapter"
	minioAdapter "AudioShare/backend/internal/adapter/minio"
	postgresAdapter "AudioShare/backend/internal/adapter/postgres"
	redisAdapter "AudioShare/backend/internal/adapter/redis"
	"AudioShare/backend/internal/config"
	"AudioShare/backend/internal/handler"
	repositoryAggregated "AudioShare/backend/internal/repository/aggregatedRepo"
	repository "AudioShare/backend/internal/repository/interfaces"
	"AudioShare/backend/internal/router"
	"AudioShare/backend/internal/server"
	"AudioShare/backend/internal/service"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

// @title Your API
// @version 1.0
// @description This is your backend API
// @host localhost:8080
// @BasePath /api/v1

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

	minioConn := adapter.MustConnect(minioAdapter.NewMinio(cfg.Minio.Host,
		cfg.Minio.Port,
		cfg.Minio.Region,
		cfg.Minio.Endpoint,
		cfg.Minio.AccessKey,
		cfg.Minio.SecretKey,
		cfg.Minio.BucketName,
		cfg.Minio.SSLMode))
	defer minioConn.Close() // just return nil. In case you need to cleanup - update func

	postgresRepository := repository.NewPostgresRepository(postgreSQLConn)
	redisRepository := repository.NewRedisRepository(redisConn)
	minioRepository := repository.NewMinioRepository(minioConn)

	aggregatedRepository := repositoryAggregated.NewAggregatedRepository(*postgresRepository,
		*redisRepository,
		*minioRepository)

	serviceInstance := service.NewService(aggregatedRepository)
	handlerInstance := handler.NewHandler(serviceInstance)
	routerInstance := router.NewRouter(handlerInstance)

	server := server.NewServer(cfg.HTTPServer.Address,
		routerInstance.InitNewRouter(),
		cfg.HTTPServer.Timeout,
		cfg.HTTPServer.Timeout)

	server.Run()
}
