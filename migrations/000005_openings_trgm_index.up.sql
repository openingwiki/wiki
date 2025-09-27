BEGIN;

CREATE EXTENSION IF NOT EXISTS unaccent;
CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE INDEX IF NOT EXISTS idx_openings_title_trgm
    ON openings USING GIN (((lower(title))) gin_trgm_ops);

END;