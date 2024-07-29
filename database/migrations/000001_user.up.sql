CREATE TABLE "users"(
    "id" bigserial PRIMARY KEY,
    "id_municipio" bigserial NOT NULL,
    "username" varchar Not null unique ,
    "password" varchar Not null,
    "email" varchar Not null unique ,
    "super_user" boolean default(false),
    "created_at"  timestamp default(now()) not null
);