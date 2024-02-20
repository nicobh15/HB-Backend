CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "households" (
  "household_id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "household_name" varchar NOT NULL,
  "address" varchar,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "users" (
  "user_id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "first_name" varchar NOT NULL,
  "password_hash" varchar NOT NULL,
  "role" varchar NOT NULL,
  "household_id" UUID,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "inventory" (
  "item_id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "household_id" UUID NOT NULL,
  "category" varchar NOT NULL,
  "name" varchar NOT NULL,
  "quantity" integer NOT NULL,
  "expiration_date" date,
  "purchase_date" date DEFAULT (now()),
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  "location" varchar
);

CREATE TABLE "recipes" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "author" varchar NOT NULL,
  "visibility" int NOT NULL DEFAULT '0',
  "data" jsonb,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE INDEX "idx_users_username" ON "users" ("username");

CREATE INDEX "idx_users_email" ON "users" ("email");

CREATE INDEX "idx_users_householdid" ON "users" ("household_id");

CREATE INDEX "idx_inventory_householdid" ON "inventory" ("household_id");

CREATE INDEX "idx_inventory_category" ON "inventory" ("category");

CREATE INDEX "idx_inventory_name" ON "inventory" ("name");

ALTER TABLE "users" ADD FOREIGN KEY ("household_id") REFERENCES "households" ("household_id");

ALTER TABLE "inventory" ADD FOREIGN KEY ("household_id") REFERENCES "households" ("household_id");

ALTER TABLE "recipes" ADD FOREIGN KEY ("author") REFERENCES "users" ("username");
