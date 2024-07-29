-- name: InsertUser :one
INSERT INTO "users"(
                    "username",
                    "password",
                    "email",
                    "super_user",
                    "id_municipio"
) VALUES ($1, $2, $3, $4,$5) RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE "username" = $1 ;

-- name: GetUserID :one
SELECT * FROM users WHERE "id" = $1 ;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
    OFFSET $2;


-- name: UpdateUser :one
UPDATE users
SET
    password = COALESCE(sqlc.narg(password), password)
WHERE
        id = sqlc.arg(id)
RETURNING *;

-- name: UpdateToSuperUser :one
UPDATE users
SET
    "super_user" = $1
WHERE
        id = sqlc.arg(id)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM "users" WHERE id =$1;