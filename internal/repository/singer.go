package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/openingwiki/wiki/internal/model"
)

type SingerRepository interface {
	CreateSinger(ctx context.Context, title string) (*model.Singer, error)
}

type PostgresSingerRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresSingerRepository(pool *pgxpool.Pool) *PostgresSingerRepository {
	return &PostgresSingerRepository{pool: pool}
}

func (r *PostgresSingerRepository) CreateSinger(ctx context.Context, name string) (*model.Singer, error) {
	const query = `
		INSERT INTO singers (name)
		VALUES ($1, NOW())
		RETURING id, created_at
	`

	var (
		id		   int64
		created_at time.Time
	)

	if err := r.pool.QueryRow(ctx, query, name).Scan(&id, &created_at); err != nil {
		return nil, err
	}

	return &model.Singer{
		ID:		id,
		Name:	name,
		CreatedAt: created_at,
	}, nil
}
