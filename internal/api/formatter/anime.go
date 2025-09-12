package formatter

import (
	"time"

	"github.com/openingwiki/wiki/internal/model"
)

// CreateAnimeRequest represents the data required to create a new anime
type CreateAnimeRequest struct {
	Title string `json:"title" binding:"required"`
}

// AnimeResponse represents an anime response
type AnimeResponse struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

// NewAnimeResponseFromDomain converts domain model to response format
func NewAnimeResponseFromDomain(m *model.Anime) *AnimeResponse {
	return &AnimeResponse{
		ID:        m.ID,
		Title:     m.Title,
		CreatedAt: m.CreatedAt,
	}
}
