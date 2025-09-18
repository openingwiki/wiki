package repository

import (
	"context"
	"fmt"
	"time"

	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/openingwiki/wiki/internal/model"
)

type AnimeRepository interface {
	CreateAnime(ctx context.Context, title string) (*model.Anime, error)
	GetAnime(ctx context.Context, id int64) (*model.Anime, error)
}

type PostgresAnimeRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresAnimeRepository(pool *pgxpool.Pool) *PostgresAnimeRepository {
	return &PostgresAnimeRepository{pool: pool}
}

func (r *PostgresAnimeRepository) CreateAnime(ctx context.Context, title string) (*model.Anime, error) {
	const query = `
		INSERT INTO anime (title, created_at) 
		VALUES ($1, NOW()) 
		RETURNING id, created_at
	`

	var (
		id      int64
		created time.Time
	)

	if err := r.pool.QueryRow(ctx, query, title).Scan(&id, &created); err != nil {
		return nil, err
	}

	return &model.Anime{
		ID:        id,
		Title:     title,
		CreatedAt: created,
	}, nil
}

func (r *PostgresAnimeRepository) GetAnime(ctx context.Context, id int64) (*model.Anime, error) {
	const query = `
		SELECT * FROM anime WHERE ID = $1
	`
	var anime model.Anime

	err := r.pool.QueryRow(ctx, query, id).Scan(&anime.ID, &anime.Title, &anime.CreatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("opening with id %d not found", id)
		}
		return nil, fmt.Errorf("get opening by id %d: %w", id, err)
	}

	return &anime, nil

}
