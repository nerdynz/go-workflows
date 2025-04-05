CREATE TABLE IF NOT EXISTS "attributes" (
  "id" TEXT NOT NULL,
  "instance_id" TEXT NOT NULL,
  "execution_id" TEXT NOT NULL,
  "data" BYTEA NOT NULL,
  PRIMARY KEY("id", "instance_id", "execution_id")
);

-- Move activity attributes to attributes table
INSERT INTO "attributes" ("id", "instance_id", "execution_id", "data") 
SELECT "id", "instance_id", "execution_id", "attributes" FROM "activities"
ON CONFLICT DO NOTHING;

ALTER TABLE "activities" DROP COLUMN "attributes";

-- Move history attributes to attributes table
INSERT INTO "attributes" ("id", "instance_id", "execution_id", "data") 
SELECT "id", "instance_id", "execution_id", "attributes" FROM "history"
ON CONFLICT DO NOTHING;

ALTER TABLE "history" DROP COLUMN "attributes";

-- Move pending_events attributes to attributes table
INSERT INTO "attributes" ("id", "instance_id", "execution_id", "data") 
SELECT "id", "instance_id", "execution_id", "attributes" FROM "pending_events"
ON CONFLICT DO NOTHING;

ALTER TABLE "pending_events" DROP COLUMN "attributes";