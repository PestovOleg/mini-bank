CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY ,
    email varchar(256) UNIQUE NOT NULL,
    phone varchar(25) NOT NULL,
    birthday timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
    name varchar(256)  NOT NULL,
    last_name varchar(256)  NOT NULL,
    patronymic varchar(256)  NOT NULL,
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL
);
