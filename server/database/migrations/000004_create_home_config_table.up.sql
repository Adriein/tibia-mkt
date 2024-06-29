CREATE TABLE IF NOT EXISTS home_config (
    id VARCHAR(36) PRIMARY KEY, -- Unique identifier (UUID)
    user_id VARCHAR(36) NOT NULL,
    cog_id VARCHAR(36) NOT NULL,
    position INTEGER,
    columns INTEGER,
    rows INTEGER,
    created_at VARCHAR(60) NOT NULL,
    updated_at VARCHAR(60) NOT NULL,
    CONSTRAINT fk_cog_id FOREIGN KEY (cog_id) REFERENCES cog(id)
);