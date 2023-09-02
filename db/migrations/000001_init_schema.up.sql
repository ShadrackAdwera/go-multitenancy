CREATE TABLE "tenant" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "name" varchar UNIQUE NOT NULL,
  "logo" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "username" varchar(50) NOT NULL,
  "email" varchar(40) NOT NULL,
  "tenant_id" uuid NOT NULL,
  "password" varchar(60) NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "tenant_policy" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "name" varchar(20) UNIQUE NOT NULL,
  "group_id" uuid,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "permission" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "name" varchar(20) UNIQUE NOT NULL,
  "description" varchar(50) NOT NULL,
  "policy_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "profile" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "user_id" uuid NOT NULL,
  "avatar" varchar,
  "group_id" uuid
);

CREATE TABLE "tenant_group" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "name" varchar(20) UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "group_policy" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "group_id" uuid,
  "policy_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_group" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "user_id" uuid,
  "group_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX ON "users" ("id", "tenant_id");

CREATE UNIQUE INDEX ON "permission" ("id", "policy_id");

CREATE UNIQUE INDEX ON "profile" ("id", "user_id");

CREATE UNIQUE INDEX ON "group_policy" ("group_id", "policy_id");

CREATE UNIQUE INDEX ON "user_group" ("group_id", "user_id");

ALTER TABLE "users" ADD FOREIGN KEY ("tenant_id") REFERENCES "tenant" ("id");

ALTER TABLE "tenant_policy" ADD FOREIGN KEY ("group_id") REFERENCES "tenant_group" ("id");

ALTER TABLE "permission" ADD FOREIGN KEY ("policy_id") REFERENCES "tenant_policy" ("id");

ALTER TABLE "profile" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "profile" ADD FOREIGN KEY ("group_id") REFERENCES "tenant_group" ("id");

ALTER TABLE "group_policy" ADD FOREIGN KEY ("group_id") REFERENCES "tenant_group" ("id");

ALTER TABLE "group_policy" ADD FOREIGN KEY ("policy_id") REFERENCES "tenant_policy" ("id");

ALTER TABLE "user_group" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_group" ADD FOREIGN KEY ("group_id") REFERENCES "tenant_group" ("id");