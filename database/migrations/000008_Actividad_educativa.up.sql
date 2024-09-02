
CREATE TYPE ultimo_grado AS ENUM ('Primaria', 'Secundaria', 'Pre Universitario','Universitario');
CREATE TABLE "estudios_actuales"(
    id bigserial primary key,
    "id_actividad_educativa" bigint not null unique,
    "tipo_enseñansa" varchar not null,
    "centro" varchar not null,
    "especialidad_grado_o_año" varchar not null,
    "año_del_dato" date not null,
    "fecha_de_graduacion" date not null
);
CREATE TABLE "actividad_educativa"(
    id bigserial primary key,
    "id_asociado" bigint NOT NULL unique ,
    "ultimo_grado_aprobado" ultimo_grado NOT NULL

);
ALTER TABLE "actividad_educativa" ADD FOREIGN KEY  ("id_asociado") REFERENCES "asociado" ("id");
ALTER TABLE "estudios_actuales" ADD FOREIGN KEY  ("id_actividad_educativa") REFERENCES "actividad_educativa" ("id");
