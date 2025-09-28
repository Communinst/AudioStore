package adapter

import (
	minioAdapter "AudioShare/backend/internal/adapter/minio"
	postgresAdapter "AudioShare/backend/internal/adapter/postgres"
	redisAdapter "AudioShare/backend/internal/adapter/redis"
	"log"
)

type Closer interface {
	Close() error
}

func MustConnect[Connection *postgresAdapter.PostgresClient | *redisAdapter.RedisClient | *minioAdapter.MinioClient](conn Connection, err error) Connection {
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

//func MustConnect(conn Closer, err error) Closer {
//	if err != nil {
//		log.Fatal(err)
//	}
//	return conn
//}
