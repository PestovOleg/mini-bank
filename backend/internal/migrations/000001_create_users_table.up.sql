CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY ,
    username varchar(256) UNIQUE NOT NULL,
    email varchar(256) UNIQUE NOT NULL,
    name varchar(256)  NOT NULL,
    last_name varchar(256)  NOT NULL,
    patronymic varchar(256)  NOT NULL,
    password varchar(256)  NOT NULL,
    is_active bool NOT NULL,
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL
);
