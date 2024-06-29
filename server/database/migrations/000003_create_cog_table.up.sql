CREATE TABLE IF NOT EXISTS cog (
     id VARCHAR(36) PRIMARY KEY, -- Unique identifier (UUID)
     name VARCHAR(255) NOT NULL,
     link VARCHAR(255) NOT NULL,
     created_at VARCHAR(60) NOT NULL,
     updated_at VARCHAR(60) NOT NULL
);