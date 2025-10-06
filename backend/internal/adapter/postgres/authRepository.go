package postgresAdapter

import (
	"AudioShare/backend/internal/entity"
	httpError "AudioShare/backend/internal/error"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type AuthPostgresRepository struct {
	db *sqlx.DB
}

func NewAuthPostgresRepository(dbWrapper *PostgresClient) *AuthPostgresRepository {
	return &AuthPostgresRepository{
		db: dbWrapper.db,
	}
}

func (this *AuthPostgresRepository) PostOne(ctx context.Context, data *entity.User) (int64, error) {

	tx, err := this.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("auth postgres repository: post one: failed to initiate transcation.")
		return -1, httpError.New(http.StatusInternalServerError,
			"auth postgres repository: post one: failed to initiate transcation.")
	}
	defer tx.Rollback()

	var resultId int64
	query := `SELECT insert_user($1, $2, $3, $4, $5, $6)`
	err = tx.QueryRowContext(ctx,
		query,
		data.Login,
		data.Email,
		data.Password,
		data.Nickname,
		data.Registered,
		data.RoleId).Scan(&resultId)
	if err == nil {
		slog.Info("auth postgres repository: post one: insertion succeded.")
		return resultId, tx.Commit()
	}

	slog.Error(fmt.Sprintf("auth postgres repository: post one: failed insertion: %s", err.Error()))
	return -1, httpError.New(http.StatusInternalServerError,
		fmt.Sprintf("auth postgres repository: post one: failed insertion: %s", err.Error()))
}

func (this *AuthPostgresRepository) GetOneByEmail(ctx context.Context, email string) (*entity.User, error) {
	var resultData entity.User
	query := `SELECT * FROM ` + user_table + ` WHERE email = $1`
	err := this.db.GetContext(ctx, &resultData, query, email)
	if err == nil {
		slog.Info("auth postgres repository: get one: by email: obtained successfully.")
		return &resultData, nil
	}
	if errors.Is(err, sql.ErrNoRows) {
		slog.Warn("auth postgres repository: get one: by email: no user by email: %s",
			slog.String("email", email),
			slog.String("table", user_table))
		return nil, nil
	}

	slog.Error(fmt.Sprintf("auth postgres repository: get one: by email: failed to obtain: %s", err.Error()))
	return nil, httpError.New(http.StatusInternalServerError,
		fmt.Sprintf("auth postgres repository: get one: by email: failed to obtain: %s", err.Error()))
}
