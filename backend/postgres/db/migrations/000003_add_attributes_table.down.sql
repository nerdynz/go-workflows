ALTER TABLE activities ADD COLUMN attributes BYTEA;
UPDATE activities SET attributes = a.data FROM attributes a WHERE activities.activity_id = a.event_id AND activities.instance_id = a.instance_id AND activities.execution_id = a.execution_id;
ALTER TABLE activities ALTER COLUMN attributes SET NOT NULL;

ALTER TABLE history ADD COLUMN attributes BYTEA;
UPDATE history SET attributes = a.data FROM attributes a WHERE history.event_id = a.event_id AND history.instance_id = a.instance_id AND history.execution_id = a.execution_id;
ALTER TABLE history ALTER COLUMN attributes SET NOT NULL;

ALTER TABLE pending_events ADD COLUMN attributes BYTEA;
UPDATE pending_events SET attributes = a.data FROM attributes a WHERE pending_events.event_id = a.event_id AND pending_events.instance_id = a.instance_id AND pending_events.execution_id = a.execution_id;
ALTER TABLE pending_events ALTER COLUMN attributes SET NOT NULL;

-- Drop attributes table
DROP TABLE attributes;