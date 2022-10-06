CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "username" varchar NOT NULL UNIQUE,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "mobile" varchar NOT NULL,
  "created_at" timestamptz DEFAULT 'now()'
);