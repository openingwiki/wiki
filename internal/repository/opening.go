package repository

import (
	"context"
	"fmt"
	"time"

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
        INSERT INTO opening (anime_id, singer_id, type, title, order_number, created_at)
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
		animeID,  // Используем переданные параметры
		singerID, // а не неинициализированные переменные
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
		return nil, fmt.Errorf("failed to create opening: %w", err)
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
