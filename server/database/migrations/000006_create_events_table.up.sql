CREATE TABLE IF NOT EXISTS events (
    id UUID PRIMARY KEY, -- Unique identifier (UUID)
    name TEXT NOT NULL,
    good_name TEXT NOT NULL,
    world TEXT NOT NULL,
    description TEXT NOT NULL,
    occurred_at TIMESTAMP(0) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_events_world_good_occurred ON prices(world, good_name, occurred_at);