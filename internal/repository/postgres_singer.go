package repository

import (
    "context"
    "errors"
    "time"

    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/openingwiki/wiki/internal/model"
)

type PostgresSingerRepo struct { pool *pgxpool.Pool }

func NewPostgresSinger(pool *pgxpool.Pool) *PostgresSingerRepo { return &PostgresSingerRepo{pool: pool} }

func (p *PostgresSingerRepo) CreateSinger(name string) (*model.Singer, error) {
    const q = `insert into singers (name, created_at) values ($1, now()) returning id, created_at`
    ctx := context.Background()
    var id int64
    var created time.Time
    if err := p.pool.QueryRow(ctx, q, name).Scan(&id, &created); err != nil { return nil, err }
    return &model.Singer{ID: id, Name: name, CreatedAt: created}, nil
}
func (p *PostgresSingerRepo) GetSinger(id int64) (*model.Singer, error) {
    const q = `select id, name, created_at from singers where id=$1`
    ctx := context.Background()
    var s model.Singer
    if err := p.pool.QueryRow(ctx, q, id).Scan(&s.ID, &s.Name, &s.CreatedAt); err != nil {
        if errors.Is(err, pgx.ErrNoRows) { return nil, errNotFound }
        return nil, err
    }
    return &s, nil
}


