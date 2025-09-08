package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/openingwiki/wiki/internal/api/formatter"
	"github.com/openingwiki/wiki/internal/service"
)

type AnimeHandler struct {
	service *service.AnimeService
}

func NewAnimeHandler(s *service.AnimeService) *AnimeHandler {
	return &AnimeHandler{service: s}
}

func (h *AnimeHandler) Register(r *gin.RouterGroup) {
	r.POST("/anime", h.createAnime)
}

func (h *AnimeHandler) createAnime(c *gin.Context) {
	var req formatter.CreateAnimeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	anime, err := h.service.CreateAnime(c.Request.Context(), req.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, formatter.NewAnimeResponseFromDomain(anime))
}
