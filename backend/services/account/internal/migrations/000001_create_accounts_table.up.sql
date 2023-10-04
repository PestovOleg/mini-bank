CREATE TABLE IF NOT EXISTS accounts (
    id uuid PRIMARY KEY ,
    user_id uuid NOT NULL,
    account varchar(20) NOT NULL,
    currency varchar(3) NOT NULL,
    name varchar(256) NOT NULL,
    amount DECIMAL(10,2) NOT NULL DEFAULT 0.00,
    interest_rate DECIMAL(5,4) NOT NULL DEFAULT 0.00,
    is_active bool NOT NULL,
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL
);

