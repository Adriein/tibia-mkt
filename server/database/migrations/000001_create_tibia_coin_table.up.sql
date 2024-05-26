CREATE TABLE IF NOT EXISTS tibia_coin (
    id VARCHAR(36) PRIMARY KEY, -- Unique identifier (UUID)
    world VARCHAR(20) NOT NULL,
    date VARCHAR(60) NOT NULL,
    buy_price DECIMAL(15,2) NOT NULL,
    sell_price DECIMAL(15,2) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_world ON tibia_coin(world);
CREATE INDEX IF NOT EXISTS idx_date ON tibia_coin(date);
CREATE INDEX IF NOT EXISTS idx_buy_price ON tibia_coin(buy_price);
CREATE INDEX IF NOT EXISTS idx_sell_price ON tibia_coin(sell_price);