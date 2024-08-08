CREATE TABLE IF NOT EXISTS kill_statistic_cron (
    id VARCHAR(36) PRIMARY KEY, -- Unique identifier (UUID)
    creature_name VARCHAR(255) NOT NULL,
    amount_killed INTEGER NOT NULL,
    drop_rate FLOAT NOT NULL,
    executed_by VARCHAR(255) NOT NULL,
    created_at VARCHAR(60) NOT NULL,
    updated_at VARCHAR(60) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_ksc_creature_name ON kill_statistic_cron(creature_name);
CREATE INDEX IF NOT EXISTS idx_ksc_created_at ON kill_statistic_cron(created_at);