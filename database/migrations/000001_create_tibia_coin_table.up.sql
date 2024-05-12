CREATE TABLE IF NOT EXISTS tibia_coin (
    id VARCHAR(36) PRIMARY KEY, -- Unique identifier (UUID)
    world VARCHAR(20) NOT NULL,
    date VARCHAR(60) NOT NULL,
    price DECIMAL(15,2) NOT NULL,
    action_type VARCHAR(3) -- Restrict values to 'BUY' or 'SELL'
);

CREATE INDEX IF NOT EXISTS idx_world ON tibia_coin(world);
CREATE INDEX IF NOT EXISTS idx_date ON tibia_coin(date);
CREATE INDEX IF NOT EXISTS idx_price ON tibia_coin(price);