package postgres

import (
	"AudioShare/backend/internal/entity"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	postgres *sqlx.DB
}

func NewUserRepository(conn *sqlx.DB) *UserRepository {
	return &UserRepository{postgres: conn}
}

func (u *UserRepository) PostOne(ctx context.Context, data *entity.User) (int, error) {
	tx, err := u.postgres.BeginTx(ctx, nil)
	if err != nil {
		return -1, fmt.Errorf("user repository: failed to post user: %w", err)
	}
	defer tx.Rollback()

	return 0, nil
}

// GetOneById(ctx context.Context, id int) (*E, error)
// GetManyById(ctx context.Context, id []int) ([]*E, er
// DeleteOneById(ctx context.Context, id int) error
// DeleteManyById(ctx context.Context, id []int) error
