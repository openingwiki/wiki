package service

import (
	"context"
	"errors"

	"github.com/openingwiki/wiki/internal/model"
	"github.com/openingwiki/wiki/internal/repository"
)

var ErrAnimeExists = errors.New("anime already exists")

type AnimeService struct {
	repo repository.AnimeRepository
}

func NewAnimeService(repo repository.AnimeRepository) *AnimeService {
	return &AnimeService{
		repo: repo,
	}
}

func (s *AnimeService) CreateAnime(ctx context.Context, title string) (*model.Anime, error) {
	return s.repo.CreateAnime(ctx, title)
}

func (s *AnimeService) GetAnime(ctx context.Context, id int64) (*model.Anime, error) {
	return s.repo.GetAnime(ctx, id)
}
