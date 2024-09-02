CREATE TABLE "asociado"(
    "id" BIGSERIAL PRIMARY KEY ,
    "nombre" VARCHAR NOT NULL ,
    "apellido1" VARCHAR NOT NULL,
    "apellido2" VARCHAR NOT NULL,
    "activo" BOOL NOT NULL,
    "carnet" BIGINT NOT NULL,
    "sexo" bool NOT NULL,
    "numeroT" BIGINT ,
    "numeroPerteneciente" VARCHAR ,
    "direccion" VARCHAR NOT NULL,
    "id_municipio" BIGINT NOT NULL
);
ALTER TABLE "asociado" ADD FOREIGN KEY  ("id_municipio") REFERENCES "municipio" ("id");