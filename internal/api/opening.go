package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/openingwiki/wiki/internal/api/formatter"
	"github.com/openingwiki/wiki/internal/service"
)

type OpeningHandler struct {
	service *service.OpeningService
}

func NewOpeningHandler(s *service.OpeningService) *OpeningHandler {
	return &OpeningHandler{service: s}
}

func (h *OpeningHandler) Register(r *gin.RouterGroup) {
	r.POST("/opening", h.createOpening)
}

func (h *OpeningHandler) createOpening(c *gin.Context) {
	var req formatter.CreateOpeningRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	opening, err := h.service.CreateOpening(c.Request.Context(), req.AnimeID, req.SingerID, req.Type, req.Title, req.OrderNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}
	c.JSON(http.StatusCreated, formatter.CreateOpeningResponseFromDomain(opening))
}
