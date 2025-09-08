package repository

import (
    "context"
    "errors"
    "time"

    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/openingwiki/wiki/internal/model"
)

type PostgresOpeningRepo struct { pool *pgxpool.Pool }

func NewPostgresOpening(pool *pgxpool.Pool) *PostgresOpeningRepo { return &PostgresOpeningRepo{pool: pool} }

func (p *PostgresOpeningRepo) CreateOpening(o *model.Opening) (*model.Opening, error) {
    const q = `insert into openings (anime_id, singer_id, type, title, order_number, created_at) values ($1,$2,$3,$4,$5, now()) returning id, created_at`
    ctx := context.Background()
    var id int64
    var created time.Time
    if err := p.pool.QueryRow(ctx, q, o.AnimeID, o.SingerID, string(o.Type), o.Title, o.OrderNumber).Scan(&id, &created); err != nil { return nil, err }
    o.ID = id
    o.CreatedAt = created
    return o, nil
}
func (p *PostgresOpeningRepo) GetOpening(id int64) (*model.Opening, error) {
    const q = `select id, anime_id, singer_id, type, title, order_number, created_at from openings where id=$1`
    ctx := context.Background()
    var o model.Opening
    var t string
    if err := p.pool.QueryRow(ctx, q, id).Scan(&o.ID, &o.AnimeID, &o.SingerID, &t, &o.Title, &o.OrderNumber, &o.CreatedAt); err != nil {
        if errors.Is(err, pgx.ErrNoRows) { return nil, errNotFound }
        return nil, err
    }
    o.Type = model.Type(t)
    return &o, nil
}


