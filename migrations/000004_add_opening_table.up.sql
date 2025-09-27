BEGIN;

CREATE EXTENSION IF NOT EXISTS unaccent;
CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE TABLE IF NOT EXISTS openings (
        id          BIGSERIAL PRIMARY KEY,
        anime_id INTEGER NOT NULL REFERENCES anime(id) ON DELETE CASCADE,
        singer_id INTEGER NOT NULL REFERENCES singers(id) ON DELETE CASCADE,
        type   TEXT NOT NULL,
        title       TEXT NOT NULL,
        order_number  INTEGER,
        created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


END;