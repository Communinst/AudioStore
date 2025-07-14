package adapter

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	DB *sqlx.DB
}

func NewPostgres(host,
	port,
	username,
	password,
	dbName,
	sslMode string) (*Postgres, func(), error) {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host,
		port,
		username,
		password,
		dbName,
		sslMode)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := sqlx.ConnectContext(ctx, "postgres", connString)
	if err != nil {
		return nil, nil, fmt.Errorf("postgres: connection failed: %w", err)
	}

	return &Postgres{DB: conn}, func() {
		if err := conn.Close(); err != nil {
			slog.Error("postgres: failed to shutdown postgres connection.", "Error:\n", err)
		}
	}, nil
}

func (p *Postgres) Close() error {
	return p.DB.Close()
}
