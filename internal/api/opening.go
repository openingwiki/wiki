package api

import (
	"github.com/gin-gonic/gin"
	"github.com/openingwiki/wiki/internal/api/formatter"
	"github.com/openingwiki/wiki/internal/service"
	"net/http"
	"strconv"
)

type OpeningHandler struct {
	service *service.OpeningService
}

func NewOpeningHandler(s *service.OpeningService) *OpeningHandler {
	return &OpeningHandler{service: s}
}

func (h *OpeningHandler) Register(r *gin.RouterGroup) {
	openingGroup := r.Group("/openings")
	openingGroup.POST("", h.createOpening)
	openingGroup.GET("/:id", h.GetOpeningByID)
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
		return
	}

	c.JSON(http.StatusCreated, formatter.CreateOpeningResponseFromDomain(opening))
}

func (h *OpeningHandler) GetOpeningByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}

	opening, err := h.service.GetOpeningByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, formatter.CreateOpeningResponseFromDomain(opening))
}
