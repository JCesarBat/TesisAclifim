-- name: CreateActividadDeportiva :one
INSERT INTO "actividad_deportiva"(
                                  "id_asociado",
                                  "aficcion_o_practica"
) VALUES ($1,$2) returning *;

-- name: CreateParticipacionD :one
INSERT INTO "participacionD"(
                              "id_actividad_deportiva",
                              "deporte",
                              "fecha",
                              "lugar_alcanzado",
                              "donde_se_desarrollo"
) VALUES ($1,$2,$3,$4,$5) returning *;

-- name: GetActividadDeportiva :one
SELECT * FROM "actividad_deportiva"  WHERE "actividad_deportiva".id =$1 ;

-- name: GetParticipacionD :many
SELECT *FROM "participacionD" WHERE "id_actividad_deportiva" =$1 ;

-- name: UpdateActividadDeportiva :one
UPDATE "actividad_deportiva" SET "aficcion_o_practica" =array_append("aficcion_o_practica",$2)
WHERE "id_asociado"=$1 returning *;

-- name: UpdateParticipacionD :one
UPDATE "participacionD"
SET
    "deporte"= $3,
    "fecha"=$4,
    "lugar_alcanzado"=$5,
    "donde_se_desarrollo"=$6
WHERE "id"=$1 and "id_actividad_deportiva"=$2  returning *;