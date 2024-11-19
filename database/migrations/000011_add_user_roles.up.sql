CREATE TABLE "rol"(
    id bigserial PRIMARY KEY ,
    rol varchar not null
);


CREATE TABLE "user_roles"(
    id bigserial PRIMARY KEY,
    user_id bigint not null ,
    rol_id bigint not null
);
ALTER TABLE "user_roles" ADD FOREIGN KEY  ("user_id") REFERENCES "users" ("id");
ALTER TABLE "user_roles" ADD FOREIGN KEY  ("rol_id") REFERENCES "rol" ("id");