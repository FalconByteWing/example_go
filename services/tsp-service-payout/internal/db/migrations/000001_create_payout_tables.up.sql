CREATE TYPE payout_status AS ENUM ('pending', 'processing', 'completed', 'failed');

CREATE TABLE IF NOT EXISTS payouts (
    id BIGSERIAL PRIMARY KEY,
    order_id VARCHAR(255) UNIQUE NOT NULL,
    bank_name VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    amount DECIMAL(20,2) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    ifsc_code VARCHAR(20) NOT NULL,
    account_no VARCHAR(50) NOT NULL,
    name VARCHAR(255) NOT NULL,
    notify_url TEXT NOT NULL,
    status payout_status NOT NULL DEFAULT 'pending',
    merchant_id UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_payouts_order_id ON payouts(order_id);
CREATE INDEX idx_payouts_merchant_id ON payouts(merchant_id);

CREATE TABLE IF NOT EXISTS merchant_balances (
    id BIGSERIAL PRIMARY KEY,
    merchant_id UUID UNIQUE NOT NULL,
    balance DECIMAL(20,2) NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_merchant_balances_merchant_id ON merchant_balances(merchant_id); 