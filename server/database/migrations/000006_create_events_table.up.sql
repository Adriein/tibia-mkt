CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    good_name TEXT NOT NULL,
    world TEXT NOT NULL,
    description TEXT NOT NULL,
    occurred_at TIMESTAMP(0) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_events_world_good_occurred ON events(world, good_name, occurred_at);