CREATE TABLE IF NOT EXISTS authentications (
    id uuid PRIMARY KEY,
    username varchar(256) UNIQUE NOT NULL,
    password_hash varchar(256) NOT NULL,
    is_active bool NOT NULL,
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL
);
