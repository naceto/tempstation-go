package sqlc

import (
	"log/slog"

	"github.com/naceto/tempstation/internal/generated/db"
	storageInterface "github.com/naceto/tempstation/internal/storage"
)

type storage struct {
	log *slog.Logger
	db  db.Querier
}

func NewStorage(db db.Querier, log *slog.Logger) storageInterface.Storage {
	return &storage{
		db:  db,
		log: log,
	}
}
