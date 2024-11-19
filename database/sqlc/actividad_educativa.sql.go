// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: actividad_educativa.sql

package database

import (
	"context"
	"time"
)

const createActividadEducativa = `-- name: CreateActividadEducativa :one
INSERT INTO "actividad_educativa" (
                                id_asociado,
                                id_estudios_actuales,
                                ultimo_grado_aprobado

) VALUES ($1,$2,$3) RETURNING id, id_asociado, id_estudios_actuales, ultimo_grado_aprobado
`

type CreateActividadEducativaParams struct {
	IDAsociado          int64       `json:"id_asociado"`
	IDEstudiosActuales  int64       `json:"id_estudios_actuales"`
	UltimoGradoAprobado UltimoGrado `json:"ultimo_grado_aprobado"`
}

func (q *Queries) CreateActividadEducativa(ctx context.Context, arg CreateActividadEducativaParams) (ActividadEducativa, error) {
	row := q.db.QueryRowContext(ctx, createActividadEducativa, arg.IDAsociado, arg.IDEstudiosActuales, arg.UltimoGradoAprobado)
	var i ActividadEducativa
	err := row.Scan(
		&i.ID,
		&i.IDAsociado,
		&i.IDEstudiosActuales,
		&i.UltimoGradoAprobado,
	)
	return i, err
}

const createEstudiosActuales = `-- name: CreateEstudiosActuales :one
INSERT INTO "estudios_actuales" (

    tipo_enseñansa,
    centro,
    especialidad_grado_o_año,
    año_del_dato,
    fecha_de_graduacion
) VALUES ($1,$2,$3,$4,$5) RETURNING id, "tipo_enseñansa", centro, "especialidad_grado_o_año", "año_del_dato", fecha_de_graduacion
`

type CreateEstudiosActualesParams struct {
	TipoEnseñansa         string    `json:"tipo_enseñansa"`
	Centro                string    `json:"centro"`
	EspecialidadGradoOAño string    `json:"especialidad_grado_o_año"`
	AñoDelDato            time.Time `json:"año_del_dato"`
	FechaDeGraduacion     time.Time `json:"fecha_de_graduacion"`
}

func (q *Queries) CreateEstudiosActuales(ctx context.Context, arg CreateEstudiosActualesParams) (EstudiosActuale, error) {
	row := q.db.QueryRowContext(ctx, createEstudiosActuales,
		arg.TipoEnseñansa,
		arg.Centro,
		arg.EspecialidadGradoOAño,
		arg.AñoDelDato,
		arg.FechaDeGraduacion,
	)
	var i EstudiosActuale
	err := row.Scan(
		&i.ID,
		&i.TipoEnseñansa,
		&i.Centro,
		&i.EspecialidadGradoOAño,
		&i.AñoDelDato,
		&i.FechaDeGraduacion,
	)
	return i, err
}

const deleteActividadEducativa = `-- name: DeleteActividadEducativa :exec
DELETE FROM "actividad_educativa" WHERE id =$1
`

func (q *Queries) DeleteActividadEducativa(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteActividadEducativa, id)
	return err
}

const deleteEstudiosActuales = `-- name: DeleteEstudiosActuales :exec
DELETE FROM "estudios_actuales" WHERE id=$1
`

func (q *Queries) DeleteEstudiosActuales(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteEstudiosActuales, id)
	return err
}

const getActividadEducativa = `-- name: GetActividadEducativa :one
SELECT id, id_asociado, id_estudios_actuales, ultimo_grado_aprobado FROM "actividad_educativa"  WHERE "actividad_educativa".id =$1
`

func (q *Queries) GetActividadEducativa(ctx context.Context, id int64) (ActividadEducativa, error) {
	row := q.db.QueryRowContext(ctx, getActividadEducativa, id)
	var i ActividadEducativa
	err := row.Scan(
		&i.ID,
		&i.IDAsociado,
		&i.IDEstudiosActuales,
		&i.UltimoGradoAprobado,
	)
	return i, err
}

const getEstudiosActuales = `-- name: GetEstudiosActuales :one
SELECT id, "tipo_enseñansa", centro, "especialidad_grado_o_año", "año_del_dato", fecha_de_graduacionFROM "estudios_actuales" WHERE "id"=$1
`

func (q *Queries) GetEstudiosActuales(ctx context.Context, id int64) (EstudiosActuale, error) {
	row := q.db.QueryRowContext(ctx, getEstudiosActuales, id)
	var i EstudiosActuale
	err := row.Scan(
		&i.ID,
		&i.TipoEnseñansa,
		&i.Centro,
		&i.EspecialidadGradoOAño,
		&i.AñoDelDato,
		&i.FechaDeGraduacion,
	)
	return i, err
}

const updateActividadEducativa = `-- name: UpdateActividadEducativa :one
UPDATE "actividad_educativa"
SET
    ultimo_grado_aprobado =$2
WHERE id =$1 returning id, id_asociado, id_estudios_actuales, ultimo_grado_aprobado
`

type UpdateActividadEducativaParams struct {
	ID                  int64       `json:"id"`
	UltimoGradoAprobado UltimoGrado `json:"ultimo_grado_aprobado"`
}

func (q *Queries) UpdateActividadEducativa(ctx context.Context, arg UpdateActividadEducativaParams) (ActividadEducativa, error) {
	row := q.db.QueryRowContext(ctx, updateActividadEducativa, arg.ID, arg.UltimoGradoAprobado)
	var i ActividadEducativa
	err := row.Scan(
		&i.ID,
		&i.IDAsociado,
		&i.IDEstudiosActuales,
		&i.UltimoGradoAprobado,
	)
	return i, err
}

const updateEstudiosActuales = `-- name: UpdateEstudiosActuales :one
UPDATE  "estudios_actuales"
SET
    tipo_enseñansa=$2,
    centro=$3,
    especialidad_grado_o_año=$4,
    año_del_dato=$5,
    fecha_de_graduacion=$6
WHERE id=$1 RETURNING id, "tipo_enseñansa", centro, "especialidad_grado_o_año", "año_del_dato", fecha_de_graduacion
`

type UpdateEstudiosActualesParams struct {
	ID                    int64     `json:"id"`
	TipoEnseñansa         string    `json:"tipo_enseñansa"`
	Centro                string    `json:"centro"`
	EspecialidadGradoOAño string    `json:"especialidad_grado_o_año"`
	AñoDelDato            time.Time `json:"año_del_dato"`
	FechaDeGraduacion     time.Time `json:"fecha_de_graduacion"`
}

func (q *Queries) UpdateEstudiosActuales(ctx context.Context, arg UpdateEstudiosActualesParams) (EstudiosActuale, error) {
	row := q.db.QueryRowContext(ctx, updateEstudiosActuales,
		arg.ID,
		arg.TipoEnseñansa,
		arg.Centro,
		arg.EspecialidadGradoOAño,
		arg.AñoDelDato,
		arg.FechaDeGraduacion,
	)
	var i EstudiosActuale
	err := row.Scan(
		&i.ID,
		&i.TipoEnseñansa,
		&i.Centro,
		&i.EspecialidadGradoOAño,
		&i.AñoDelDato,
		&i.FechaDeGraduacion,
	)
	return i, err
}