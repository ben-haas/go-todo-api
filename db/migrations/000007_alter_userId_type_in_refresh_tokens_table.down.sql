ALTER TABLE refresh_tokens
    ALTER COLUMN user_id SET DATA TYPE INT USING user_id::INT;