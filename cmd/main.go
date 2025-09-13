package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/openingwiki/wiki/docs"
	"github.com/openingwiki/wiki/internal/api"
	"github.com/openingwiki/wiki/internal/config"
	"github.com/openingwiki/wiki/internal/repository"
	"github.com/openingwiki/wiki/internal/service"
)

// @title           Swagger OpeningWiki
// @version         1.0
// @termsOfService  http://swagger.io/terms/

// @license.name  MIT LICENSE
// @license.url   https://github.com/openingwiki/wiki/blob/main/LICENSE

// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

func main() {
	cfg := config.Load()

	var pool *pgxpool.Pool
	var err error

	if cfg.DatabaseURL != "" {
		pool, err = pgxpool.New(context.Background(), cfg.DatabaseURL)
		if err != nil {
			log.Fatalf("failed to create db pool: %v", err)
		}
		defer pool.Close()

		// Ping database to ensure connection is alive
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := pool.Ping(ctx); err != nil {
			log.Fatalf("failed to ping database: %v", err)
		}

		log.Println("Connected to database")
	} else {
		log.Println("No DATABASE_URL provided")
	}

	// Initialize repository and service
	animeRepo := repository.NewPostgresAnimeRepository(pool)
	animeService := service.NewAnimeService(animeRepo)
	openingRepo := repository.NewPostgresOpeningRepository(pool)
	openingService := service.NewOpeningService(openingRepo)

	// Initialize Gin and register routes
	r := gin.Default()
	api.NewRouter(r, animeService, openingService)
	// Start server
	log.Printf("Starting server on %s...", cfg.HTTPAddr)
	if err := r.Run(cfg.HTTPAddr); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
