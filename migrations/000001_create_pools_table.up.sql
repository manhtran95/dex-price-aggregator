-- Create pools table
CREATE TABLE IF NOT EXISTS pools (
    id SERIAL PRIMARY KEY,
    address VARCHAR(42) NOT NULL UNIQUE,
    dex VARCHAR(50) NOT NULL,
    token0 VARCHAR(42) NOT NULL,
    token1 VARCHAR(42) NOT NULL,
    fee INTEGER DEFAULT 3000,
    created_block BIGINT,
    is_active BOOLEAN DEFAULT true,
    last_synced_block BIGINT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Indexes for fast lookups
CREATE INDEX idx_pools_tokens ON pools(token0, token1);
CREATE INDEX idx_pools_dex ON pools(dex);
CREATE INDEX idx_pools_active ON pools(is_active);

-- Store token metadata
CREATE TABLE IF NOT EXISTS tokens (
    address VARCHAR(42) PRIMARY KEY,
    symbol VARCHAR(20),
    name VARCHAR(100),
    decimals INTEGER,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Track sync progress
CREATE TABLE IF NOT EXISTS sync_status (
    id SERIAL PRIMARY KEY,
    dex VARCHAR(50) NOT NULL,
    factory_address VARCHAR(42) NOT NULL,
    last_synced_block BIGINT DEFAULT 0,
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(dex, factory_address)
);