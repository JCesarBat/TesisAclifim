package database

import (
	"context"
	"database/sql"
)

type Store interface {
	Querier
	DinamicList(ctx context.Context, param DinamicListParam) ([]Asociado, error)
}

type SQLStore struct {
	conn *sql.DB
	*Queries
}

// NewStore creates a new store
func NewStore(conn *sql.DB) Store {

	return &SQLStore{
		conn:    conn,
		Queries: New(conn),
	}
}
