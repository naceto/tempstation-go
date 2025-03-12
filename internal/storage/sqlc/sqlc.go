package sqlc

import (
	"errors"
	"log/slog"

	"github.com/naceto/tempstation/internal/generated/db"
	storageInterface "github.com/naceto/tempstation/internal/storage"
	"github.com/naceto/tempstation/internal/storage/sqlc/convert"
)

var (
	FormatStorageError = "storage error: %w"
	ErrDuplicateEntry  = errors.New("duplicate entry")
	ErrInternal        = errors.New("internal server error")
)

type storage struct {
	db      db.Querier
	convert convert.Convert
	log     *slog.Logger
}

func NewStorage(db db.Querier, convert convert.Convert, log *slog.Logger) storageInterface.Storage {
	return &storage{
		db:      db,
		convert: convert,
		log:     log,
	}
}
