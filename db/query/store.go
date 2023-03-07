package query

import (
	"github.com/jmoiron/sqlx"
)

// Store provides all function to execute db queries and transactions
type Store struct {
	db *sqlx.DB
}

// NewStore creates a new store
func NewStore(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}
