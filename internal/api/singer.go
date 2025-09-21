package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/openingwiki/wiki/internal/api/formatter"
	"github.com/openingwiki/wiki/internal/service"
)

type SingerHandler struct {
	service *service.SingerService
}

func NewSingerHandler(s *service.SingerService) *SingerHandler {
	return &SingerHandler{service: s}
}

func (h *SingerHandler) Register(r *gin.RouterGroup) {
	r.POST("/singers", h.createSinger)
}

func (h *SingerHandler) createSinger(c *gin.Context) {
	var req formatter.CreateSingerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	singer, err := h.service.CreateSinger(c.Request.Context(), req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, formatter.NewSingerResponseFromDomain(singer))
}