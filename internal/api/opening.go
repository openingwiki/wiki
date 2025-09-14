package api

import (
	"net/http"
	"strconv"

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
	openingGroup := r.Group("/openings")
	openingGroup.POST("", h.createOpening)
	openingGroup.GET("/:id", h.GetOpeningByID)
}

// CreateOpening godoc
// @Summary Create a new opening
// @Description Add a new anime opening theme to the database
// @Tags openings
// @Accept json
// @Produce json
// @Param request body formatter.CreateOpeningRequest true "Opening creation data"
// @Success 201 {object} formatter.OpeningResponse "Successfully created opening"
// @Failure 400 {object} map[string]interface{} "Invalid input data"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /openings [post]
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

// GetOpeningByID godoc
// @Summary Get opening by ID
// @Description Get detailed information about a specific opening by its ID
// @Tags openings
// @Produce json
// @Param id path int true "Opening ID"
// @Success 200 {object} formatter.OpeningResponse "Successfully retrieved opening"
// @Failure 400 {object} map[string]interface{} "get opening by id"
// @Failure 404 {object} map[string]interface{} "opening with id %d not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /openings/{id} [get]
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
