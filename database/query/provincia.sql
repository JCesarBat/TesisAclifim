-- name: GetProvincia :one
SELECT * FROM "provincia" WHERE id =$1;

-- name: GetAllProv :many
SELECT * FROM "provincia" ;


-- name: InertarProv :one
Insert INTO "provincia" ("name") VALUES ($1) RETURNING *;