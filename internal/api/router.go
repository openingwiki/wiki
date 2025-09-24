package api

import (
	"github.com/gin-gonic/gin"
	"github.com/openingwiki/wiki/internal/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewRouter registers all API routes under /api/v1
func NewRouter(
	r *gin.Engine,
	animeService *service.AnimeService,
	openingService *service.OpeningService,
	singerService *service.SingerService,
) {
	v1 := r.Group("/api/v1")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	NewAnimeHandler(animeService).Register(v1)
	NewOpeningHandler(openingService).Register(v1)
	NewSingerHandler(singerService).Register(v1)
}
