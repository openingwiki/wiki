BEGIN;

CREATE TABLE IF NOT EXISTS anime (
     id          BIGSERIAL PRIMARY KEY,
     title       TEXT NOT NULL,
     created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

END;
