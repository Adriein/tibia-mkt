CREATE TABLE IF NOT EXISTS good (
     id VARCHAR(36) PRIMARY KEY, -- Unique identifier (UUID)
     name VARCHAR(255) NOT NULL,
     link VARCHAR(255) NOT NULL,
     creatures JSONB NOT NULL,
     created_at VARCHAR(60) NOT NULL,
     updated_at VARCHAR(60) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_g_name ON good(name);