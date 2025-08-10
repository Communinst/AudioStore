package adapter

import (
	"AudioShare/backend/internal/adapter/minio"
	"AudioShare/backend/internal/adapter/postgres"
	"AudioShare/backend/internal/adapter/redis"
	"log"
)

type Closer interface {
	Close() error
}

func MustConnect[Connection *Postgres | *Redis | *Minio](conn Connection, err error) Connection {
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
