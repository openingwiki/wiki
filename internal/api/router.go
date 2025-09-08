package api

import (
	"github.com/gin-gonic/gin"
	"github.com/openingwiki/wiki/internal/service"
)

// NewRouter registers all API routes under /api/v1
func NewRouter(
	r *gin.Engine,
	animeService *service.AnimeService,
) {
	v1 := r.Group("/api/v1")

	NewAnimeHandler(animeService).Register(v1)
}
