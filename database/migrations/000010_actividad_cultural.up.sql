CREATE TABLE "participacionC"(
   "id" bigserial primary key ,
   "id_actividad_cultural" bigint,
    "especialidad" varchar not null,
    "fecha" date not null,
    "lugar_alcanzado" int,
    "donde_se_desarrollo" text not null

);
CREATE TABLE "actividad_cultural"(
    "id" bigserial primary key ,
    "id_asociado" bigint not null unique,
    "especialidad" varchar[]



);
ALTER TABLE "actividad_cultural" ADD FOREIGN KEY  ("id_asociado") REFERENCES "asociado" ("id");
ALTER TABLE "participacionC" ADD FOREIGN KEY  ("id_actividad_cultural") REFERENCES "actividad_cultural" ("id");