package imagerepository

import "github.com/jmoiron/sqlx"

type imageStore struct {
	db *sqlx.DB
}

func NewImageStore(db *sqlx.DB) *imageStore {
	return &imageStore{db: db}
}
