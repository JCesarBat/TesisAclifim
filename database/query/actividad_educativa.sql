-- name: CreateActividadEducativa :one
INSERT INTO "actividad_educativa" (
                                id_asociado,
                                id_estudios_actuales,
                                ultimo_grado_aprobado

) VALUES ($1,$2,$3) RETURNING *;

-- name: CreateEstudiosActuales :one
INSERT INTO "estudios_actuales" (

    tipo_enseñansa,
    centro,
    especialidad_grado_o_año,
    año_del_dato,
    fecha_de_graduacion
) VALUES ($1,$2,$3,$4,$5) RETURNING *;

-- name: GetActividadEducativa :one
SELECT * FROM "actividad_educativa"  WHERE "actividad_educativa".id =$1 ;

-- name: GetEstudiosActuales :one
SELECT *FROM "estudios_actuales" WHERE "id"=$1 ;

-- name: UpdateActividadEducativa :one
UPDATE "actividad_educativa"
SET
    ultimo_grado_aprobado =$2
WHERE id =$1 returning *;

-- name: UpdateEstudiosActuales :one
UPDATE  "estudios_actuales"
SET
    tipo_enseñansa=$2,
    centro=$3,
    especialidad_grado_o_año=$4,
    año_del_dato=$5,
    fecha_de_graduacion=$6
WHERE id=$1 RETURNING *;

-- name: DeleteActividadEducativa :exec
DELETE FROM "actividad_educativa" WHERE id =$1;

-- name: DeleteEstudiosActuales :exec
DELETE FROM "estudios_actuales" WHERE id=$1;