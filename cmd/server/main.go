package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/openingwiki/wiki/internal/api"
	"github.com/openingwiki/wiki/internal/config"
	"github.com/openingwiki/wiki/internal/repository"
	"github.com/openingwiki/wiki/internal/service"
)

func main() {
	cfg := config.Load()

	var repos repository.Repos
	if cfg.DatabaseURL != "" {
		pool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}
		repos = repository.Repos{
			Anime:   repository.NewPostgresAnime(pool),
			Singer:  repository.NewPostgresSinger(pool),
			Opening: repository.NewPostgresOpening(pool),
		}
		log.Println("using PostgreSQL repository")
	} else {
		mem := repository.NewInMemory()
		repos = repository.Repos{Anime: mem, Singer: mem, Opening: mem}
		log.Println("using in-memory repository")
	}

	svc := service.New(repos)
	h := api.Router(svc)

	log.Println("wiki service listening on", cfg.HTTPAddr)
	if err := http.ListenAndServe(cfg.HTTPAddr, h); err != nil {
		log.Fatal(err)
	}
}