CREATE TABLE IF NOT EXISTS prices (
    id UUID PRIMARY KEY,
    good_name VARCHAR(255) NOT NULL,
    world VARCHAR(20) NOT NULL,
    buy_price DECIMAL(15,2) NOT NULL,
    sell_price DECIMAL(15,2) NOT NULL,
    created_at VARCHAR(60) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_prices_good_name ON prices(good_name);
CREATE INDEX IF NOT EXISTS idx_prices_world ON prices(world);
CREATE INDEX IF NOT EXISTS idx_prices_buy_price ON prices(buy_price);
CREATE INDEX IF NOT EXISTS idx_prices_sell_price ON prices(sell_price);
CREATE INDEX IF NOT EXISTS idx_prices_created_at ON prices(created_at);

CREATE INDEX idx_prices_world_good_created ON prices(world, good_name, created_at ASC);