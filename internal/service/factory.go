package service

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/naceto/tempstation/configs"
	"github.com/naceto/tempstation/internal/generated/db"
)

type DB interface {
	db.DBTX
	Close() error
}

type DependencyFactory interface {
	GetDB(context.Context, *configs.Config) (DB, error)
}

type factory struct{}

func NewDependencyFactory() DependencyFactory {
	return &factory{}
}

func (d *factory) GetDB(ctx context.Context, cfg *configs.Config) (DB, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.PostgresUsername, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDBName)
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
