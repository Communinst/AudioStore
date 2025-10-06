package postgresAdapter

import (
	"AudioShare/backend/internal/entity"
	"context"

	"github.com/jmoiron/sqlx"
)

type DumpPostgresRepository struct {
	db *sqlx.DB
}

func NewDumpPostgresRepository(dbWrapper *PostgresClient) *DumpPostgresRepository {
	return &DumpPostgresRepository{db: dbWrapper.db}
}

func (r *DumpPostgresRepository) InsertDump(ctx context.Context, dump *entity.Dump) error {
	query := "INSERT INTO dumps (filename, size) VALUES ($1, $2)"
	_, err := r.db.ExecContext(ctx, query, dump.Filename, dump.Size)
	return err
}

func (r *DumpPostgresRepository) GetAllDumps(ctx context.Context) ([]entity.Dump, error) {
	var dumps []entity.Dump
	query := "SELECT * FROM dumps"
	err := r.db.SelectContext(ctx, &dumps, query)
	return dumps, err
}
