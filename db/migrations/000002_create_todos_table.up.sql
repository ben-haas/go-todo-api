CREATE TABLE IF NOT EXISTS todos
(
    id          BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    title       TEXT      NOT NULL,
    description TEXT      NOT NULL,
    priority    TEXT      NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    user_id     BIGINT REFERENCES users (id) ON DELETE CASCADE
);
