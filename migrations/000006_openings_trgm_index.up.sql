CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_openings_title_trgm
    ON openings USING GIN ((lower(public.immutable_unaccent(title))) gin_trgm_ops);