package api

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/openingwiki/wiki/internal/service"
)

func RegisterAnimeRoutes(r gin.IRoutes, s *service.Service) {
    r.POST("/anime", handleCreateAnime(s))
    r.GET("/anime/:id", handleGetAnime(s))
}

func handleCreateAnime(s *service.Service) gin.HandlerFunc {
    type body struct{ Title string `json:"title"` }
    return func(c *gin.Context) {
        var b body
        if err := c.BindJSON(&b); err != nil { c.String(http.StatusBadRequest, err.Error()); return }
        a, err := s.CreateAnime(b.Title)
        if err != nil { c.String(http.StatusBadRequest, err.Error()); return }
        c.JSON(http.StatusCreated, a)
    }
}

func handleGetAnime(s *service.Service) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
        a, err := s.GetAnime(id)
        if err != nil { c.String(http.StatusNotFound, "not found"); return }
        c.JSON(http.StatusOK, a)
    }
}


