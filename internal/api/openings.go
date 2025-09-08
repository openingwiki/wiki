package api

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/openingwiki/wiki/internal/model"
    "github.com/openingwiki/wiki/internal/service"
)

func RegisterOpeningRoutes(r gin.IRoutes, s *service.Service) {
    r.POST("/openings", handleCreateOpening(s))
    r.GET("/openings/:id", handleGetOpening(s))
}

func handleCreateOpening(s *service.Service) gin.HandlerFunc {
    type body struct {
        AnimeID     int64      `json:"anime_id"`
        SingerID    int64      `json:"singer_id"`
        Type        model.Type `json:"type"`
        Title       string     `json:"title"`
        OrderNumber int        `json:"order_number"`
    }
    return func(c *gin.Context) {
        var b body
        if err := c.BindJSON(&b); err != nil { c.String(http.StatusBadRequest, err.Error()); return }
        o, err := s.CreateOpening(b.AnimeID, b.SingerID, b.Type, b.Title, b.OrderNumber)
        if err != nil { c.String(http.StatusBadRequest, err.Error()); return }
        c.JSON(http.StatusCreated, o)
    }
}

func handleGetOpening(s *service.Service) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
        o, err := s.GetOpening(id)
        if err != nil { c.String(http.StatusNotFound, "not found"); return }
        c.JSON(http.StatusOK, o)
    }
}


