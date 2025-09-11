BEGIN;
CREATE TABLE IF NOT EXISTS singers (
    id          BIGSERIAL PRIMARY KEY,
    name        TEXT,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
END;