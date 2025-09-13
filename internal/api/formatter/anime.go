package formatter

import (
	"time"

	"github.com/openingwiki/wiki/internal/model"
)

type CreateAnimeRequest struct {
	Title string `json:"title" binding:"required"`
}

type AnimeResponse struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

func NewAnimeResponseFromDomain(m *model.Anime) *AnimeResponse {
	return &AnimeResponse{
		ID:        m.ID,
		Title:     m.Title,
		CreatedAt: m.CreatedAt,
	}
}
