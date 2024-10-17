package repository

import "github.com/jmoiron/sqlx"

type blogMarketStore struct {
	db *sqlx.DB
}

func NewBlogMarketStore(db *sqlx.DB) *blogMarketStore {
	return &blogMarketStore{db: db}
}
