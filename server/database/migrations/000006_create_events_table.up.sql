CREATE TABLE IF NOT EXISTS events (
    id UUID PRIMARY KEY, -- Unique identifier (UUID)
    name TEXT NOT NULL,
    good_name TEXT NOT NULL,
    world TEXT NOT NULL,
    description TEXT NOT NULL,
    occurred_at TIMESTAMP(0) NOT NULL
);