package userrepository

import "github.com/jmoiron/sqlx"

type userStore struct {
	db *sqlx.DB
}

func NewUserStore(db *sqlx.DB) *userStore {
	return &userStore{db: db}
}
