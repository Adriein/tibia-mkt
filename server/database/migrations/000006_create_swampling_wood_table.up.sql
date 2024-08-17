CREATE TABLE IF NOT EXISTS swampling_wood (
    id VARCHAR(36) PRIMARY KEY, -- Unique identifier (UUID)
    world VARCHAR(20) NOT NULL,
    date VARCHAR(60) NOT NULL,
    buy_price DECIMAL(15,2) NOT NULL,
    sell_price DECIMAL(15,2) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_sw_world ON swampling_wood(world);
CREATE INDEX IF NOT EXISTS idx_sw_date ON swampling_wood(date);
CREATE INDEX IF NOT EXISTS idx_sw_buy_price ON swampling_wood(buy_price);
CREATE INDEX IF NOT EXISTS idx_sw_sell_price ON swampling_wood(sell_price);