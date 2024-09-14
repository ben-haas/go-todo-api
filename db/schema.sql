CREATE TABLE IF NOT EXISTS todos
(
    id          INTEGER PRIMARY KEY,
    title       TEXT                                                 NOT NULL,
    description TEXT                                                 NOT NULL,
    priority    TEXT CHECK ( priority IN ('LOW', 'MEDIUM', 'HIGH') ) NOT NULL DEFAULT 'LOW',
    complete    BOOLEAN                                              NOT NULL DEFAULT FALSE,
    user_id     INTEGER                                              NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS users
(
    id         INTEGER PRIMARY KEY,
    email      TEXT    NOT NULL,
    password   TEXT    NOT NULL,
    created_at TEXT    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_admin   BOOLEAN NOT NULL DEFAULT FALSE
);