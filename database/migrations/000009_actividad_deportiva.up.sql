CREATE TABLE "participacionD"(
    "id" bigserial primary key ,
    "id_actividad_deportiva" bigint,
    "deporte" varchar not null,
    "fecha" date not null,
    "lugar_alcanzado" int,
    "donde_se_desarrollo" text not null

);
CREATE TABLE "actividad_deportiva"(
    "id" bigserial primary key ,
    "id_asociado" bigint not null unique,
    "aficcion_o_practica" varchar[]


);
ALTER TABLE "actividad_deportiva" ADD FOREIGN KEY  ("id_asociado") REFERENCES "asociado" ("id");
ALTER TABLE "participacionD" ADD FOREIGN KEY  ("id_actividad_deportiva") REFERENCES "actividad_deportiva" ("id");