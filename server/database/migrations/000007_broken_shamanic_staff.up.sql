CREATE TABLE IF NOT EXISTS broken_shamanic_staff (
    id VARCHAR(36) PRIMARY KEY, -- Unique identifier (UUID)
    world VARCHAR(20) NOT NULL,
    date VARCHAR(60) NOT NULL,
    buy_price DECIMAL(15,2) NOT NULL,
    sell_price DECIMAL(15,2) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_bss_world ON broken_shamanic_staff(world);
CREATE INDEX IF NOT EXISTS idx_bss_date ON broken_shamanic_staff(date);
CREATE INDEX IF NOT EXISTS idx_bss_buy_price ON broken_shamanic_staff(buy_price);
CREATE INDEX IF NOT EXISTS idx_bss_sell_price ON broken_shamanic_staff(sell_price);