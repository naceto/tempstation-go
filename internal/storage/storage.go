package storage

import "github.com/naceto/tempstation/internal/generated/db"

type Storage interface{}

type storage struct {
	db db.Querier
}

func NewStorage(db db.Querier) Storage {
	return &storage{
		db: db,
	}
}
