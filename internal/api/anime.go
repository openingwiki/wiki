package api

import (
	"net/http"
	"strconv"

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
	r.GET("/anime/:id", h.getAnime)
}

// createAnime godoc
// @Summary Create a new anime
// @Description Add a new anime to the database
// @Tags anime
// @Accept json
// @Produce json
// @Param request body formatter.CreateAnimeRequest true "Anime creation data"
// @Success 201 {object} formatter.AnimeResponse "Successfully created anime"
// @Failure 400 {object} map[string]interface{} "Invalid input data"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /anime [post]
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

// getAnime godoc
// @Summary Get anime by id
// @Description Get anime by id from the database
// @Tags anime
// @Accept json
// @Produce json
// @Param id path int true "Anime id"
// @Success 200 {object} formatter.AnimeResponse "Successfully retrieved anime"
// @Failure 400 {object} map[string]interface{} "Invalid anime id"
// @Failure 404 {object} map[string]interface{} "Anime not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /anime/{id} [get]
func (h *AnimeHandler) getAnime(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid anime id"})
		return
	}

	anime, err := h.service.GetAnime(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, formatter.NewAnimeResponseFromDomain(anime))
}
