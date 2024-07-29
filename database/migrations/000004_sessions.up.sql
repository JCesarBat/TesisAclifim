CREATE TABLE "sessions"(
                           "id" uuid PRIMARY KEY ,
                           "user_id" bigint NOT NULL,
                           "refresh_token" VARCHAR NOT NULL,
                           "user_agent" VARCHAR NOT NULL,
                           "client_ip" VARCHAR NOT NULL,
                           "is_blocked" boolean NOT NULL DEFAULT false,
                           "expires_at" TIMESTAMPTZ NOT NULL,
                           "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

ALTER TABLE "sessions" ADD FOREIGN KEY  ("user_id") REFERENCES "users" ("id");