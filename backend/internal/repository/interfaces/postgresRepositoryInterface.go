package repository

import (
	postgresAdapter "AudioShare/backend/internal/adapter/postgres"
	"AudioShare/backend/internal/entity"
	"context"
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"golang.org/x/crypto/bcrypt"
)

type AuthPostgresRepositoryInterface interface {
	PostOne(ctx context.Context, data *entity.User) (int64, error)
	GetOneByEmail(ctx context.Context, email string) (*entity.User, error)
}

type DumpPostgresRepositoryInterface interface {
	InsertDump(ctx context.Context, dump *entity.Dump) error
	GetAllDumps(ctx context.Context) ([]entity.Dump, error)
}

type EntityPostgresRepository[E Entity] interface {
	PostOne(ctx context.Context, data *E) (int64, error)
	GetOneById(ctx context.Context, id uint64) (*E, error)
	GetAll(ctx context.Context) ([]*E, error)
	DeleteOneById(ctx context.Context, id uint64) error
}

type UserPostgresRepositoryInterface interface {
	EntityPostgresRepository[entity.User]
	CheckIfUserWithRoleExists(ctx context.Context, roleId uint8) (bool, error)
}

type PostgresRepository struct {
	Auth AuthPostgresRepositoryInterface
	User UserPostgresRepositoryInterface
	Dump DumpPostgresRepositoryInterface
}

func NewPostgresRepository(dbWrapper *postgresAdapter.PostgresClient) *PostgresRepository {
	postgresRepository := &PostgresRepository{
		Auth: postgresAdapter.NewAuthPostgresRepository(dbWrapper),
		User: postgresAdapter.NewUserPostgresRepository(dbWrapper),
		Dump: postgresAdapter.NewDumpPostgresRepository(dbWrapper),
	}

	if err := postgresRepository.InitFirstAdmin(); err != nil {
		log.Fatal(err)
	}

	return postgresRepository
}

func (this *PostgresRepository) InitFirstAdmin() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var AdminCreds struct {
		Login      string `env:"ADMIN_LOGIN" env-default:"admin"`
		Email      string `env:"ADMIN_EMAIL" env-default:"admin@admin.com"`
		Password   string `env:"ADMIN_PASSWORD" env-required:"true"`
		Nickname   string `env:"ADMIN_NICKNAME" env-default:"ASCENDED"`
		Registered time.Time
		RoleId     uint8 `env:"ADMIN_DEFAULT_ROLE" env-required:"true"`
	}
	// Loaded via respecting .env
	err := cleanenv.ReadEnv(&AdminCreds)
	//fmt.Printf("%v\n", AdminCreds)
	if err != nil {
		slog.Error("Failed to initialise fisrt admin.")
		return fmt.Errorf("Failed to initialise first admin.")
	}

	// Pass get hashed
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(AdminCreds.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// Pass changed
	AdminCreds.Password = string(hashedPassword)

	// Verify admin absence
	exists, err := this.User.CheckIfUserWithRoleExists(ctx, AdminCreds.RoleId)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	// Build user and call tx
	_, err = this.Auth.PostOne(ctx, &entity.User{
		Login:      AdminCreds.Login,
		Email:      AdminCreds.Email,
		Password:   AdminCreds.Password,
		Nickname:   AdminCreds.Nickname,
		Registered: time.Now(),
		RoleId:     AdminCreds.RoleId,
	})
	if err != nil {
		return err
	}
	slog.Info("Admin added.")
	return nil
}
