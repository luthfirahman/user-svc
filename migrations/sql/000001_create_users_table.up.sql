CREATE TABLE "users" (
    "id" serial PRIMARY KEY,
    "name" VARCHAR(100) NOT NULL,
    "phone_number" VARCHAR(14) UNIQUE NOT NULL,
    "password" VARCHAR(250) NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMP NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("phone_number");