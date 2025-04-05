CREATE TABLE IF NOT EXISTS instances (
  id BIGSERIAL PRIMARY KEY,
  instance_id TEXT NOT NULL,
  execution_id TEXT NOT NULL,
  parent_instance_id TEXT,
  parent_execution_id TEXT,
  parent_schedule_event_id BIGINT,
  metadata BYTEA,
  state INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  completed_at TIMESTAMP,
  locked_until TIMESTAMP,
  sticky_until TIMESTAMP,
  worker TEXT,

  UNIQUE (instance_id, execution_id)
);

CREATE INDEX idx_instances_locked_until_completed_at ON instances (completed_at, locked_until, sticky_until, worker);
CREATE INDEX idx_instances_parent_instance_id_parent_execution_id ON instances (parent_instance_id, parent_execution_id);

CREATE TABLE IF NOT EXISTS pending_events (
  id BIGSERIAL PRIMARY KEY,
  event_id TEXT NOT NULL,
  sequence_id BIGINT NOT NULL,
  instance_id TEXT NOT NULL,
  execution_id TEXT NOT NULL,
  event_type INTEGER NOT NULL,
  timestamp TIMESTAMP NOT NULL,
  schedule_event_id BIGINT NOT NULL,
  attributes BYTEA NOT NULL,
  visible_at TIMESTAMP,

  INDEX idx_pending_events_inid_exid (instance_id, execution_id),
  INDEX idx_pending_events_inid_exid_visible_at_schedule_event_id (instance_id, execution_id, visible_at, schedule_event_id)
);

CREATE TABLE IF NOT EXISTS history (
  id BIGSERIAL PRIMARY KEY,
  event_id TEXT NOT NULL,
  sequence_id BIGINT NOT NULL,
  instance_id TEXT NOT NULL,
  execution_id TEXT NOT NULL,
  event_type INTEGER NOT NULL,
  timestamp TIMESTAMP NOT NULL,
  schedule_event_id BIGINT NOT NULL,
  attributes BYTEA NOT NULL,
  visible_at TIMESTAMP,

  INDEX idx_history_instance_id_execution_id (instance_id, execution_id),
  INDEX idx_history_instance_id_execution_id_sequence_id (instance_id, execution_id, sequence_id)
);

CREATE TABLE IF NOT EXISTS activities (
  id BIGSERIAL PRIMARY KEY,
  activity_id TEXT NOT NULL,
  instance_id TEXT NOT NULL,
  execution_id TEXT NOT NULL,
  event_type INTEGER NOT NULL,
  timestamp TIMESTAMP NOT NULL,
  schedule_event_id BIGINT NOT NULL,
  attributes BYTEA NOT NULL,
  visible_at TIMESTAMP,
  locked_until TIMESTAMP,
  worker TEXT,

  UNIQUE (instance_id, execution_id, activity_id, worker)
);

CREATE INDEX idx_activities_locked_until ON activities (locked_until);