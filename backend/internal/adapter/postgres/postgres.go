package postgresAdapter

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	user_table = "users"
)



//TODO

// Repos initiation
// Service peeked

type PostgresClient struct {
	db *sqlx.DB
}

func NewPostgres(host,
	port,
	username,
	password,
	dbName,
	sslMode string) (*PostgresClient, error) {
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
	return &PostgresClient{db: conn}, nil
}

func (p *PostgresClient) Close() error {
	if err := p.db.Close(); err != nil {
		return fmt.Errorf("postgres: failed to shutdown postgres connection: %w", err)
	}
	return nil
}
