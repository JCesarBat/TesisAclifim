-- name: GetMunicipio :one
SELECT * FROM "municipio" WHERE id =$1;

-- name: GetAllMunicipio :many
SELECT * FROM "municipio" WHERE id_provincia =$1;

-- name: InsertMunicipio :one
INSERT IN9TO "municipio"("name","id_provincia")values ($1,$2) RETURNING *;