package formatter

import (
	"time"

	"github.com/openingwiki/wiki/internal/model"
)

type CreateOpeningRequest struct {
	ID          int64             `json:"id"`
	AnimeID     int64             `json:"anime_id" binding:"required"`
	SingerID    int64             `json:"singer_id" binding:"required"`
	Type        model.OpeningType `json:"type" binding:"required"`
	Title       string            `json:"title"`
	OrderNumber int64             `json:"order_number"`
	CreatedAt   time.Time
}

type OpeningResponse struct {
	ID          int64             `json:"id"`
	AnimeID     int64             `json:"anime_id" binding:"required"`
	SingerID    int64             `json:"singer_id" binding:"required"`
	Type        model.OpeningType `json:"type" binding:"required"`
	Title       string            `json:"title"`
	OrderNumber int64             `json:"order_number"`
	CreatedAt   time.Time
}

func CreateOpeningResponseFromDomain(m *model.Opening) *OpeningResponse {
	return &OpeningResponse{
		ID:          m.ID,
		AnimeID:     m.AnimeId,
		SingerID:    m.SingerId,
		Type:        m.Type,
		Title:       m.Title,
		OrderNumber: m.OrderNumber,
		CreatedAt:   m.CreatedAt,
	}

}
