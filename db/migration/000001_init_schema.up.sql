CREATE TABLE "episodes" (
  "id" bigserial PRIMARY KEY,
  "number" integer UNIQUE NOT NULL,
  "anime" bigserial NOT NULL,
  "description" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "animes" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "description" varchar NOT NULL,
  "status" varchar NOT NULL,
  "num_of_episodes" integer,
  "cast" varchar[] NOT NULL,
  "genre" varchar[] NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "characters" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "from_anime" bigserial NOT NULL,
  "intro" varchar NOT NULL,
  "characteristics" varchar[] NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "episodes" ADD FOREIGN KEY ("anime") REFERENCES "animes" ("id");

ALTER TABLE "characters" ADD FOREIGN KEY ("from_anime") REFERENCES "animes" ("id");

CREATE INDEX ON "episodes" ("number");

CREATE INDEX ON "animes" ("name");

CREATE INDEX ON "animes" ("status");

CREATE INDEX ON "characters" ("name");

CREATE INDEX ON "characters" ("name", "from_anime");
