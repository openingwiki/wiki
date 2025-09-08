package repository

import (
    "context"
    "errors"
    "time"

    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/openingwiki/wiki/internal/model"
)

type PostgresAnimeRepo struct { pool *pgxpool.Pool }

func NewPostgresAnime(pool *pgxpool.Pool) *PostgresAnimeRepo { return &PostgresAnimeRepo{pool: pool} }

func (p *PostgresAnimeRepo) CreateAnime(title string) (*model.Anime, error) {
    const q = `insert into anime (title, created_at) values ($1, now()) returning id, created_at`
    ctx := context.Background()
    var id int64
    var created time.Time
    if err := p.pool.QueryRow(ctx, q, title).Scan(&id, &created); err != nil { return nil, err }
    return &model.Anime{ID: id, Title: title, CreatedAt: created}, nil
}
func (p *PostgresAnimeRepo) GetAnime(id int64) (*model.Anime, error) {
    const q = `select id, title, created_at from anime where id=$1`
    ctx := context.Background()
    var a model.Anime
    if err := p.pool.QueryRow(ctx, q, id).Scan(&a.ID, &a.Title, &a.CreatedAt); err != nil {
        if errors.Is(err, pgx.ErrNoRows) { return nil, errNotFound }
        return nil, err
    }
    return &a, nil
}


