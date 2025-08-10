package adapter

import (
	"log"
)

type Closer interface {
	Close() error
}

func MustConnect[Connection *postgres.Postgres | *redis.Redis | *minio.Minio](conn Connection, err error) Connection {
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
