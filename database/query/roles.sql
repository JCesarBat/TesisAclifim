-- name: CreateRol :one
INSERT INTO "rol" (
                   rol
) VALUES ($1) RETURNING *;

-- name: InsertRolUser :one
INSERT INTO "user_roles" (
                          user_id,
                          rol_id
) VALUES ($1,$2) RETURNING *;

-- name: GetAllRol :many
SELECT * FROM "rol" ;

-- name: GetAllRolFromUser :many
SELECT * FROM "rol" WHERE "rol".id = (SELECT id FROM "user_roles" WHERE user_id = $1);