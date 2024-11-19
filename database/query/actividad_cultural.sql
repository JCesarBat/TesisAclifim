-- name: CreateActividadCultural :one
INSERT INTO "actividad_cultural" ("id_asociado",
                                  "especialidad")
VALUES ($1,$2) RETURNING *;

-- name: CreateParticipacionCultural :one
INSERT INTO "participacionC" (
    "id_actividad_cultural",
    "especialidad",
    "fecha",
    "lugar_alcanzado",
    "donde_se_desarrollo"
) VALUES ($1,$2,$3,$4,$5) RETURNING *;

-- name: GetActividadCultural :one
SELECT * FROM "actividad_cultural" WHERE id=$1;

-- name: GetParticipacionC :many
SELECT * FROM "participacionC" WHERE id_actividad_cultural = $1 ;

-- name: UpdateParticipacionC :one
UPDATE "participacionC" SET
        "especialidad" =$2,
        "fecha"=$3,
        "lugar_alcanzado"=$4,
        "donde_se_desarrollo"=$5
    WHERE id =$1 RETURNING *;

-- name: UpdateActividadCultural :one
UPDATE "actividad_cultural" SET especialidad =array_append("especialidad",$2)
 WHERE id = $1  RETURNING *;

-- name: DeleteActividadCulural :exec
DELETE FROM "actividad_cultural" WHERE id =$1;

-- name: DeleteParticipacionC :exec
DELETE FROM "participacionC" WHERE id =$1;
