-- name: InsertAsoiciado :one
INSERT INTO "asociado"("nombre", "apellido1", "apellido2","activo","carnet","sexo","numeroT","numeroPerteneciente","direccion","id_municipio")
VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) returning *;

-- name: GetAsociado :one
SELECT * FROM "asociado" WHERE id = $1;

-- name: ListAsociado :many
SELECT * FROM "asociado"
ORDER BY id
LIMIT $1
    OFFSET $2;

-- name: UpdateAsociado :one
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
WHERE id = $1  returning *;

-- name: DeleteAsociado :exec
DELETE FROM "asociado" WHERE id = $1;