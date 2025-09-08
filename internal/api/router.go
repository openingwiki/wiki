package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/openingwiki/wiki/internal/service"
)

// Router composes all domain routes under /api/v1 using gin
func Router(s *service.Service) http.Handler {
	r := gin.New()
	api := r.Group("/api/v1")
	RegisterAnimeRoutes(api, s)
	RegisterSingerRoutes(api, s)
	RegisterOpeningRoutes(api, s)
	return r
}
