CREATE TABLE "municipio"(
                "id" bigserial PRIMARY KEY,
                "id_provincia" bigserial NOT NULL,
                "name" varchar(255) NOT NULL UNIQUE ,
                FOREIGN KEY ("id_provincia") REFERENCES provincia("id")
);