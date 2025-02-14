CREATE TABLE IF NOT EXISTS users
(
    id         BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    email      TEXT      NOT NULL,
    password   TEXT      NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);