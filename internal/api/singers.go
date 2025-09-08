package api

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/openingwiki/wiki/internal/service"
)

func RegisterSingerRoutes(r gin.IRoutes, s *service.Service) {
    r.POST("/singers", handleCreateSinger(s))
    r.GET("/singers/:id", handleGetSinger(s))
}

func handleCreateSinger(s *service.Service) gin.HandlerFunc {
    type body struct{ Name string `json:"name"` }
    return func(c *gin.Context) {
        var b body
        if err := c.BindJSON(&b); err != nil { c.String(http.StatusBadRequest, err.Error()); return }
        sg, err := s.CreateSinger(b.Name)
        if err != nil { c.String(http.StatusBadRequest, err.Error()); return }
        c.JSON(http.StatusCreated, sg)
    }
}

func handleGetSinger(s *service.Service) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
        sng, err := s.GetSinger(id)
        if err != nil { c.String(http.StatusNotFound, "not found"); return }
        c.JSON(http.StatusOK, sng)
    }
}


