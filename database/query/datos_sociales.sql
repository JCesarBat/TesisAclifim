-- name: GetDatosSociales :one
SELECT * FROM "datos_sociales" WHERE "id_asociado" =$1;

-- name: InsertDatosSociales :one
INSERT INTO "datos_sociales" ("id_asociado",
                              "ocupacion",
                              "estado_civil",
                              "integracion_revolucionaria")VALUES($1,$2,$3,$4)returning *;


-- name: UpdateDatosSociales :one
UPDATE  "datos_sociales" SET
                "ocupacion"=$2,
                "estado_civil"=$3,
                "integracion_revolucionaria"=$4 WHERE
                "id_asociado"=$1 returning *;

-- name: DeleteDatosSociales :exec
DELETE FROM "datos_sociales" WHERE "id_asociado" =$1;