package marketrepository

import "github.com/jmoiron/sqlx"

type marketStore struct {
	db *sqlx.DB
}

func NewMarketStore(db *sqlx.DB) *marketStore {
	return &marketStore{db: db}
}
