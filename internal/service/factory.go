package service

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/naceto/tempstation/configs"
	"github.com/naceto/tempstation/internal/generated/db"
)

type DB interface {
	db.DBTX
	Close(context.Context) error
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
	db, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
