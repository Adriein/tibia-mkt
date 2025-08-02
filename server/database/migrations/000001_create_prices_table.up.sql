CREATE TABLE IF NOT EXISTS prices (
    id UUID PRIMARY KEY,
    offer_type VARCHAR(4) NOT NULL,
    good_name VARCHAR(255) NOT NULL,
    world VARCHAR(20) NOT NULL,
    created_by VARCHAR(30) NOT NULL,
    good_amount SMALLINT NOT NULL,
    unit_price BIGINT NOT NULL,
    total_price BIGINT NOT NULL,
    end_at TIMESTAMP(2) NOT NULL,
    created_at TIMESTAMP(2) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_prices_good_name ON prices(good_name);
CREATE INDEX IF NOT EXISTS idx_prices_world ON prices(world);
CREATE INDEX IF NOT EXISTS idx_prices_unit_price ON prices(unit_price);
CREATE INDEX IF NOT EXISTS idx_prices_end_at ON prices(end_at);
CREATE INDEX IF NOT EXISTS idx_prices_created_at ON prices(created_at);

CREATE INDEX IF NOT EXISTS idx_prices_world_good_created ON prices(world, good_name, created_at, end_at);