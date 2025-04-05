CREATE TABLE IF NOT EXISTS "instances" (
  "id" TEXT NOT NULL,
  "execution_id" TEXT NOT NULL,
  "parent_instance_id" TEXT NULL,
  "parent_execution_id" TEXT NULL,
  "parent_schedule_event_id" BIGINT NULL,
  "metadata" TEXT NULL,
  "state" BIGINT NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "completed_at" TIMESTAMP NULL,
  "locked_until" TIMESTAMP NULL,
  "sticky_until" TIMESTAMP NULL,
  "worker" TEXT NULL,
  PRIMARY KEY("id", "execution_id")
);

CREATE INDEX IF NOT EXISTS "idx_instances_id_execution_id" ON "instances" ("id", "execution_id");
CREATE INDEX IF NOT EXISTS "idx_instances_locked_until_completed_at" ON "instances" ("locked_until", "sticky_until", "completed_at", "worker");
CREATE INDEX IF NOT EXISTS "idx_instances_parent_instance_id_parent_execution_id" ON "instances" ("parent_instance_id", "parent_execution_id");

CREATE TABLE IF NOT EXISTS "pending_events" (
  "id" TEXT,
  "sequence_id" BIGINT NOT NULL, -- not used but keep for now for query compat
  "instance_id" TEXT NOT NULL,
  "execution_id" TEXT NOT NULL,
  "event_type" BIGINT NOT NULL,
  "timestamp" TIMESTAMP NOT NULL,
  "schedule_event_id" BIGINT NOT NULL,
  "attributes" BYTEA NOT NULL,
  "visible_at" TIMESTAMP NULL,
  PRIMARY KEY("id", "instance_id")
);

CREATE INDEX IF NOT EXISTS "idx_pending_events_instance_id_execution_id_visible_at_schedule_event_id" ON "pending_events" ("instance_id", "execution_id", "visible_at", "schedule_event_id");

CREATE TABLE IF NOT EXISTS "history" (
  "id" TEXT,
  "sequence_id" BIGINT NOT NULL,
  "instance_id" TEXT NOT NULL,
  "execution_id" TEXT NOT NULL,
  "event_type" BIGINT NOT NULL,
  "timestamp" TIMESTAMP NOT NULL,
  "schedule_event_id" BIGINT NOT NULL,
  "attributes" BYTEA NOT NULL,
  "visible_at" TIMESTAMP NULL,
  PRIMARY KEY("id", "instance_id")
);

CREATE INDEX IF NOT EXISTS "idx_history_instance_sequence_id" ON "history" ("instance_id", "execution_id", "sequence_id");

CREATE TABLE IF NOT EXISTS "activities" (
  "id" TEXT PRIMARY KEY,
  "instance_id" TEXT NOT NULL,
  "execution_id" TEXT NOT NULL,
  "event_type" BIGINT NOT NULL,
  "timestamp" TIMESTAMP NOT NULL,
  "schedule_event_id" BIGINT NOT NULL,
  "attributes" BYTEA NOT NULL,
  "visible_at" TIMESTAMP NULL,
  "locked_until" TIMESTAMP NULL,
  "worker" TEXT NULL
);

CREATE INDEX IF NOT EXISTS "idx_activities_id_worker" ON "activities" ("id", "worker");
CREATE INDEX IF NOT EXISTS "idx_activities_locked_until" ON "activities" ("locked_until");
CREATE INDEX IF NOT EXISTS "idx_activities_instance_id_execution_id_worker" ON "activities" ("instance_id", "execution_id", "worker");
