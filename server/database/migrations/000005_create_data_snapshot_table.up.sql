CREATE TABLE IF NOT EXISTS data_snapshot_cron (
    id VARCHAR(36) PRIMARY KEY, -- Unique identifier (UUID)
    cog VARCHAR(255) NOT NULL,
    std_deviation FLOAT NOT NULL,
    mean INT NOT NULL,
    total_droped INT NOT NULL,
    executed_by VARCHAR(255) NOT NULL,
    created_at VARCHAR(60) NOT NULL,
    updated_at VARCHAR(60) NOT NULL
    );

CREATE INDEX IF NOT EXISTS idx_dsc_cog ON data_snapshot_cron(cog);
CREATE INDEX IF NOT EXISTS idx_dsc_created_at ON data_snapshot_cron(created_at);