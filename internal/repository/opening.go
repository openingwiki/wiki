package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/openingwiki/wiki/internal/model"
)

type OpeningRepository interface {
	CreateOpening(
		ctx context.Context,
		animeID int64,
		singerID int64,
		openingType model.OpeningType,
		title string,
		orderNumber int64,
	) (*model.Opening, error)
	GetOpeningByID(ctx context.Context, openingID int64) (*model.Opening, error)
	SearchOpeningByTitle(ctx context.Context, title string) ([]model.OpeningSearchItem, error)
}

type PostgresOpeningRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresOpeningRepository(pool *pgxpool.Pool) *PostgresOpeningRepository {
	return &PostgresOpeningRepository{pool: pool}
}

func (r *PostgresOpeningRepository) CreateOpening(
	ctx context.Context,
	animeID int64,
	singerID int64,
	openingType model.OpeningType,
	title string,
	orderNumber int64,
) (*model.Opening, error) {
	const query = `
        INSERT INTO openings (anime_id, singer_id, type, title, order_number, created_at)
        VALUES ($1, $2, $3, $4, $5, NOW())
        RETURNING id, anime_id, singer_id, type, title, order_number, created_at
    `

	var (
		id        int64
		animeId   int64
		singerId  int64
		oType     model.OpeningType
		oTitle    string
		oOrder    int64
		createdAt time.Time
	)

	err := r.pool.QueryRow(
		ctx,
		query,
		animeID,
		singerID,
		openingType,
		title,
		orderNumber,
	).Scan(
		&id,
		&animeId,
		&singerId,
		&oType,
		&oTitle,
		&oOrder,
		&createdAt,
	)
	if err != nil {
		return nil, fmt.Errorf("create opening: %w", err)
	}

	return &model.Opening{
		ID:          id,
		AnimeId:     animeId,
		SingerId:    singerId,
		Type:        oType,
		Title:       oTitle,
		OrderNumber: oOrder,
		CreatedAt:   createdAt,
	}, nil
}

func (r *PostgresOpeningRepository) GetOpeningByID(ctx context.Context, id int64) (*model.Opening, error) {
	const query = `
        SELECT id, anime_id, singer_id, type, title, order_number, created_at
        FROM openings 
        WHERE id = $1`

	var (
		openingID int64
		animeId   int64
		singerId  int64
		oType     model.OpeningType
		oTitle    string
		oOrder    int64
		createdAt time.Time
	)

	err := r.pool.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&openingID,
		&animeId,
		&singerId,
		&oType,
		&oTitle,
		&oOrder,
		&createdAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("opening with id %d: not found", id)
		}
		return nil, fmt.Errorf("get opening by id %d: %w", id, err)
	}

	return &model.Opening{
		ID:          openingID,
		AnimeId:     animeId,
		SingerId:    singerId,
		Type:        oType,
		Title:       oTitle,
		OrderNumber: oOrder,
		CreatedAt:   createdAt,
	}, nil
}

func (r *PostgresOpeningRepository) SearchOpeningByTitle(ctx context.Context, title string) ([]model.OpeningSearchItem, error) {
	const query = `
	SELECT id,
           title
	FROM openings
	WHERE lower(public.unaccent(title)) % lower(public.unaccent($1))
	ORDER BY title
	LIMIT 20;
	`

	rows, err := r.pool.Query(
		ctx,
		query,
		title,
	)
	if err != nil {
		return nil, fmt.Errorf("search openings by title: %w", err)
	}
	var out []model.OpeningSearchItem

	for rows.Next() {
		var it model.OpeningSearchItem
		if err := rows.Scan(&it.ID, &it.Title); err != nil {
			rows.Close()
			return nil, fmt.Errorf("scan search row: %w", err)
		}
		out = append(out, it)
	}

	if err := rows.Err(); err != nil {
		rows.Close()
		return nil, fmt.Errorf("rows: %w", err)
	}
	rows.Close()
	return out, nil
}
