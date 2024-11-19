// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: provincia.sql

package database

import (
	"context"
)

const getAllProv = `-- name: GetAllProv :many
SELECT id, name FROM "provincia"
`

func (q *Queries) GetAllProv(ctx context.Context) ([]Provincium, error) {
	rows, err := q.db.QueryContext(ctx, getAllProv)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Provincium{}
	for rows.Next() {
		var i Provincium
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProvincia = `-- name: GetProvincia :one
SELECT id, name FROM "provincia" WHERE id =$1
`

func (q *Queries) GetProvincia(ctx context.Context, id int64) (Provincium, error) {
	row := q.db.QueryRowContext(ctx, getProvincia, id)
	var i Provincium
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const inertarProv = `-- name: InertarProv :one
Insert INTO "provincia" ("name") VALUES ($1) RETURNING id, name
`

func (q *Queries) InertarProv(ctx context.Context, name string) (Provincium, error) {
	row := q.db.QueryRowContext(ctx, inertarProv, name)
	var i Provincium
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}
