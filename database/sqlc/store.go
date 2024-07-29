package database

import "database/sql"

type Store interface {
	Querier
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
