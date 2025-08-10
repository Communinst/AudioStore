package postgres

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//TODO
// Finish Minio conn. Test it.

// Repos initiation
// Service peeked

type Postgres struct {
	db *sqlx.DB
}

func NewPostgres(host,
	port,
	username,
	password,
	dbName,
	sslMode string) (*Postgres, error) {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host,
		port,
		username,
		password,
		dbName,
		sslMode)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := sqlx.ConnectContext(ctx, dbName, connString)
	if err != nil {
		return nil, fmt.Errorf("postgres: connection establishment failed: %w", err)
	}

	slog.Info("postgres: connection established.")
	return &Postgres{db: conn}, nil
}

func (p *Postgres) Close() error {
	if err := p.db.Close(); err != nil {
		return fmt.Errorf("postgres: failed to shutdown postgres connection: %w", err)
	}
	return nil
}
