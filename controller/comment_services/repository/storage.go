package commmentrepository

import "github.com/jmoiron/sqlx"

type commentStore struct {
	db *sqlx.DB
}

func NewCommentStore(db *sqlx.DB) *commentStore {
	return &commentStore{db: db}
}
