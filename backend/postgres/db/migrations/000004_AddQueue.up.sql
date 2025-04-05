ALTER TABLE instances ADD COLUMN IF NOT EXISTS queue TEXT DEFAULT '';

-- Update index
DROP INDEX IF EXISTS idx_instances_locked_until_completed_at;
CREATE INDEX idx_instances_locked_until_completed_at_queue ON instances (completed_at, locked_until, sticky_until, worker, queue);

ALTER TABLE activities ADD COLUMN IF NOT EXISTS queue TEXT DEFAULT '';

-- Update index
DROP INDEX IF EXISTS idx_activities_locked_until;
CREATE INDEX idx_activities_locked_until_queue ON activities (locked_until, queue);