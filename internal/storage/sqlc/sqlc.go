package sqlc

import (
	"github.com/naceto/tempstation/internal/generated/db"
	storageInterface "github.com/naceto/tempstation/internal/storage"
)

type storage struct {
	db db.Querier
}

func NewStorage(db db.Querier) storageInterface.Storage {
	return &storage{
		db: db,
	}
}
