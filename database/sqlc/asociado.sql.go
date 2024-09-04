// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: asociado.sql

package database

import (
	"context"
	"database/sql"
)

const deleteAsociado = `-- name: DeleteAsociado :exec
DELETE FROM "asociado" WHERE id = $1
`

func (q *Queries) DeleteAsociado(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAsociado, id)
	return err
}

const getAsociado = `-- name: GetAsociado :one
SELECT id, nombre, apellido1, apellido2, activo, carnet, sexo, "numeroT", "numeroPerteneciente", direccion, id_municipio FROM "asociado" WHERE id = $1
`

func (q *Queries) GetAsociado(ctx context.Context, id int64) (Asociado, error) {
	row := q.db.QueryRowContext(ctx, getAsociado, id)
	var i Asociado
	err := row.Scan(
		&i.ID,
		&i.Nombre,
		&i.Apellido1,
		&i.Apellido2,
		&i.Activo,
		&i.Carnet,
		&i.Sexo,
		&i.NumeroT,
		&i.NumeroPerteneciente,
		&i.Direccion,
		&i.IDMunicipio,
	)
	return i, err
}

const insertAsoiciado = `-- name: InsertAsoiciado :one
INSERT INTO "asociado"("nombre", "apellido1", "apellido2","activo","carnet","sexo","numeroT","numeroPerteneciente","direccion","id_municipio")
VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) returning id, nombre, apellido1, apellido2, activo, carnet, sexo, "numeroT", "numeroPerteneciente", direccion, id_municipio
`

type InsertAsoiciadoParams struct {
	Nombre              string         `json:"nombre"`
	Apellido1           string         `json:"apellido1"`
	Apellido2           string         `json:"apellido2"`
	Activo              bool           `json:"activo"`
	Carnet              int64          `json:"carnet"`
	Sexo                bool           `json:"sexo"`
	NumeroT             sql.NullInt64  `json:"numeroT"`
	NumeroPerteneciente sql.NullString `json:"numeroPerteneciente"`
	Direccion           string         `json:"direccion"`
	IDMunicipio         int64          `json:"id_municipio"`
}

func (q *Queries) InsertAsoiciado(ctx context.Context, arg InsertAsoiciadoParams) (Asociado, error) {
	row := q.db.QueryRowContext(ctx, insertAsoiciado,
		arg.Nombre,
		arg.Apellido1,
		arg.Apellido2,
		arg.Activo,
		arg.Carnet,
		arg.Sexo,
		arg.NumeroT,
		arg.NumeroPerteneciente,
		arg.Direccion,
		arg.IDMunicipio,
	)
	var i Asociado
	err := row.Scan(
		&i.ID,
		&i.Nombre,
		&i.Apellido1,
		&i.Apellido2,
		&i.Activo,
		&i.Carnet,
		&i.Sexo,
		&i.NumeroT,
		&i.NumeroPerteneciente,
		&i.Direccion,
		&i.IDMunicipio,
	)
	return i, err
}

const listAsociado = `-- name: ListAsociado :many
SELECT id, nombre, apellido1, apellido2, activo, carnet, sexo, "numeroT", "numeroPerteneciente", direccion, id_municipio FROM "asociado"
ORDER BY id
LIMIT $1
    OFFSET $2
`

type ListAsociadoParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAsociado(ctx context.Context, arg ListAsociadoParams) ([]Asociado, error) {
	rows, err := q.db.QueryContext(ctx, listAsociado, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Asociado{}
	for rows.Next() {
		var i Asociado
		if err := rows.Scan(
			&i.ID,
			&i.Nombre,
			&i.Apellido1,
			&i.Apellido2,
			&i.Activo,
			&i.Carnet,
			&i.Sexo,
			&i.NumeroT,
			&i.NumeroPerteneciente,
			&i.Direccion,
			&i.IDMunicipio,
		); err != nil {
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

const updateAsociado = `-- name: UpdateAsociado :one
UPDATE "asociado"
set
    "nombre" = $2,
    "apellido1" = $3,
    "apellido2" = $4,
    "activo" = $5,
    "carnet" = $6,
    "sexo" = $7,
    "numeroT" = $8,
    "numeroPerteneciente" = $9,
    "direccion" = $10,
    "id_municipio" = $11
WHERE id = $1  returning id, nombre, apellido1, apellido2, activo, carnet, sexo, "numeroT", "numeroPerteneciente", direccion, id_municipio
`

type UpdateAsociadoParams struct {
	ID                  int64          `json:"id"`
	Nombre              string         `json:"nombre"`
	Apellido1           string         `json:"apellido1"`
	Apellido2           string         `json:"apellido2"`
	Activo              bool           `json:"activo"`
	Carnet              int64          `json:"carnet"`
	Sexo                bool           `json:"sexo"`
	NumeroT             sql.NullInt64  `json:"numeroT"`
	NumeroPerteneciente sql.NullString `json:"numeroPerteneciente"`
	Direccion           string         `json:"direccion"`
	IDMunicipio         int64          `json:"id_municipio"`
}

func (q *Queries) UpdateAsociado(ctx context.Context, arg UpdateAsociadoParams) (Asociado, error) {
	row := q.db.QueryRowContext(ctx, updateAsociado,
		arg.ID,
		arg.Nombre,
		arg.Apellido1,
		arg.Apellido2,
		arg.Activo,
		arg.Carnet,
		arg.Sexo,
		arg.NumeroT,
		arg.NumeroPerteneciente,
		arg.Direccion,
		arg.IDMunicipio,
	)
	var i Asociado
	err := row.Scan(
		&i.ID,
		&i.Nombre,
		&i.Apellido1,
		&i.Apellido2,
		&i.Activo,
		&i.Carnet,
		&i.Sexo,
		&i.NumeroT,
		&i.NumeroPerteneciente,
		&i.Direccion,
		&i.IDMunicipio,
	)
	return i, err
}