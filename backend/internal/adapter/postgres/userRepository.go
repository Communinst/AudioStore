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
	_ "github.com/lib/pq"
)

type UserPostgresRepository struct {
	db *sqlx.DB
}

func NewUserPostgresRepository(connWrapper *PostgresClient) *UserPostgresRepository {
	return &UserPostgresRepository{
		db: connWrapper.db,
	}
}

func (u *UserPostgresRepository) PostOne(ctx context.Context, data *entity.User) (int64, error) {

	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("user postgres repository: post one: failed to initiate transcation.")
		return -1, httpError.New(http.StatusInternalServerError,
			"user postgres repository: post one: failed to initiate transcation.")
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
		slog.Info("user postgres repository: post one: insertion succeded.")
		return resultId, tx.Commit()
	}

	slog.Error(fmt.Sprintf("user postgres repository: post one: failed insertion: %s", err.Error()))
	return -1, httpError.New(http.StatusInternalServerError,
		fmt.Sprintf("user postgres repository: post one: failed insertion: %s", err.Error()))
}

func (u *UserPostgresRepository) GetOneById(ctx context.Context, id uint64) (*entity.User, error) {

	var resultData *entity.User
	query := `SELECT * FROM get_user_by_id($1)`
	err := u.db.GetContext(ctx, &resultData, query, id)
	if err == nil {
		slog.Info("user postgres repository: get one: by id: obtained successfully.")
		return resultData, nil
	}
	if errors.Is(err, sql.ErrNoRows) {
		slog.Error("user postgres repository: get one: by id: no user by id: ", slog.Uint64("user_id", id))
	}

	slog.Error(fmt.Sprintf("user postgres repository: get one: by id: failed to obtain: %s", err.Error()))
	return nil, httpError.New(http.StatusInternalServerError,
		fmt.Sprintf("user postgres repository: get one: by id: failed to obtain: %s", err.Error()))
}

func (u *UserPostgresRepository) GetAll(ctx context.Context) ([]*entity.User, error) {

	var resultData []*entity.User
	query := `SELECT * FROM get_all_users()`
	err := u.db.SelectContext(ctx, &resultData, query)
	if err == nil {
		slog.Info("user repo: get all: obtained successfully.")
		return resultData, nil
	}

	return nil, httpError.New(http.StatusInternalServerError,
		fmt.Sprintf("user postgres repository: get all: failed to obtain: %s", err.Error()))
}

func (u *UserPostgresRepository) DeleteOneById(ctx context.Context, id uint64) error {

	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("user postgres repository: delete one: by id: failed to initiate transcation.")
		return httpError.New(http.StatusInternalServerError,
			"user postgres repository: delete one: by id: failed to initiate transcation.")
	}
	defer tx.Rollback()

	var affected uint64
	query := `SELECT delete_user_by_id($1)`
	err = tx.QueryRowContext(ctx, query, id).
		Scan(&affected)

	if err != nil {
		slog.Error("user postgres repository: delete one: by id: failed deletion.")
		return httpError.New(http.StatusInternalServerError,
			"user postgres repository: delete one: by id: failed deletion.")
	}
	if affected > 0 {
		slog.Info(fmt.Sprintf("user postgres repository: delete one: by id: affected rows: %d.", affected))
	} else {
		slog.Info(fmt.Sprintf("user postgres repository: delete one: by id: user bt %d id wasn't found.", affected))
	}

	return tx.Commit()
}
