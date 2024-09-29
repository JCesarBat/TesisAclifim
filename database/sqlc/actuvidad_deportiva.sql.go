// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: actuvidad_deportiva.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

const createActividadDeportiva = `-- name: CreateActividadDeportiva :one
INSERT INTO "actividad_deportiva"(
                                  "id_asociado",
                                  "aficcion_o_practica"
) VALUES ($1,$2) returning id, id_asociado, aficcion_o_practica
`

type CreateActividadDeportivaParams struct {
	IDAsociado        int64    `json:"id_asociado"`
	AficcionOPractica []string `json:"aficcion_o_practica"`
}

func (q *Queries) CreateActividadDeportiva(ctx context.Context, arg CreateActividadDeportivaParams) (ActividadDeportiva, error) {
	row := q.db.QueryRowContext(ctx, createActividadDeportiva, arg.IDAsociado, pq.Array(arg.AficcionOPractica))
	var i ActividadDeportiva
	err := row.Scan(&i.ID, &i.IDAsociado, pq.Array(&i.AficcionOPractica))
	return i, err
}

const createParticipacionD = `-- name: CreateParticipacionD :one
INSERT INTO "participacionD"(
                              "id_actividad_deportiva",
                              "deporte",
                              "fecha",
                              "lugar_alcanzado",
                              "donde_se_desarrollo"
) VALUES ($1,$2,$3,$4,$5) returning id, id_actividad_deportiva, deporte, fecha, lugar_alcanzado, donde_se_desarrollo
`

type CreateParticipacionDParams struct {
	IDActividadDeportiva sql.NullInt64 `json:"id_actividad_deportiva"`
	Deporte              string        `json:"deporte"`
	Fecha                time.Time     `json:"fecha"`
	LugarAlcanzado       sql.NullInt32 `json:"lugar_alcanzado"`
	DondeSeDesarrollo    string        `json:"donde_se_desarrollo"`
}

func (q *Queries) CreateParticipacionD(ctx context.Context, arg CreateParticipacionDParams) (ParticipacionD, error) {
	row := q.db.QueryRowContext(ctx, createParticipacionD,
		arg.IDActividadDeportiva,
		arg.Deporte,
		arg.Fecha,
		arg.LugarAlcanzado,
		arg.DondeSeDesarrollo,
	)
	var i ParticipacionD
	err := row.Scan(
		&i.ID,
		&i.IDActividadDeportiva,
		&i.Deporte,
		&i.Fecha,
		&i.LugarAlcanzado,
		&i.DondeSeDesarrollo,
	)
	return i, err
}

const getActividadDeportiva = `-- name: GetActividadDeportiva :one
SELECT id, id_asociado, aficcion_o_practica FROM "actividad_deportiva"  WHERE "actividad_deportiva".id =$1
`

func (q *Queries) GetActividadDeportiva(ctx context.Context, id int64) (ActividadDeportiva, error) {
	row := q.db.QueryRowContext(ctx, getActividadDeportiva, id)
	var i ActividadDeportiva
	err := row.Scan(&i.ID, &i.IDAsociado, pq.Array(&i.AficcionOPractica))
	return i, err
}

const getParticipacionD = `-- name: GetParticipacionD :many
SELECT id, id_actividad_deportiva, deporte, fecha, lugar_alcanzado, donde_se_desarrolloFROM "participacionD" WHERE "id_actividad_deportiva" =$1
`

func (q *Queries) GetParticipacionD(ctx context.Context, idActividadDeportiva sql.NullInt64) ([]ParticipacionD, error) {
	rows, err := q.db.QueryContext(ctx, getParticipacionD, idActividadDeportiva)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ParticipacionD{}
	for rows.Next() {
		var i ParticipacionD
		if err := rows.Scan(
			&i.ID,
			&i.IDActividadDeportiva,
			&i.Deporte,
			&i.Fecha,
			&i.LugarAlcanzado,
			&i.DondeSeDesarrollo,
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

const updateActividadDeportiva = `-- name: UpdateActividadDeportiva :one
UPDATE "actividad_deportiva" SET "aficcion_o_practica" =array_append("aficcion_o_practica",$2)
WHERE "id_asociado"=$1 returning id, id_asociado, aficcion_o_practica
`

type UpdateActividadDeportivaParams struct {
	IDAsociado  int64       `json:"id_asociado"`
	ArrayAppend interface{} `json:"array_append"`
}

func (q *Queries) UpdateActividadDeportiva(ctx context.Context, arg UpdateActividadDeportivaParams) (ActividadDeportiva, error) {
	row := q.db.QueryRowContext(ctx, updateActividadDeportiva, arg.IDAsociado, arg.ArrayAppend)
	var i ActividadDeportiva
	err := row.Scan(&i.ID, &i.IDAsociado, pq.Array(&i.AficcionOPractica))
	return i, err
}

const updateParticipacionD = `-- name: UpdateParticipacionD :one
UPDATE "participacionD"
SET
    "deporte"= $3,
    "fecha"=$4,
    "lugar_alcanzado"=$5,
    "donde_se_desarrollo"=$6
WHERE "id"=$1 and "id_actividad_deportiva"=$2  returning id, id_actividad_deportiva, deporte, fecha, lugar_alcanzado, donde_se_desarrollo
`

type UpdateParticipacionDParams struct {
	ID                   int64         `json:"id"`
	IDActividadDeportiva sql.NullInt64 `json:"id_actividad_deportiva"`
	Deporte              string        `json:"deporte"`
	Fecha                time.Time     `json:"fecha"`
	LugarAlcanzado       sql.NullInt32 `json:"lugar_alcanzado"`
	DondeSeDesarrollo    string        `json:"donde_se_desarrollo"`
}

func (q *Queries) UpdateParticipacionD(ctx context.Context, arg UpdateParticipacionDParams) (ParticipacionD, error) {
	row := q.db.QueryRowContext(ctx, updateParticipacionD,
		arg.ID,
		arg.IDActividadDeportiva,
		arg.Deporte,
		arg.Fecha,
		arg.LugarAlcanzado,
		arg.DondeSeDesarrollo,
	)
	var i ParticipacionD
	err := row.Scan(
		&i.ID,
		&i.IDActividadDeportiva,
		&i.Deporte,
		&i.Fecha,
		&i.LugarAlcanzado,
		&i.DondeSeDesarrollo,
	)
	return i, err
}
