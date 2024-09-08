CREATE TABLE IF NOT EXISTS refresh_tokens
(
    id         BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    token      TEXT      NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    user_id    INT REFERENCES users (id) ON DELETE CASCADE
);