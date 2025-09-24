package service

import (
	"context"

	"github.com/openingwiki/wiki/internal/model"
	"github.com/openingwiki/wiki/internal/repository"
)

type SingerService struct {
	repo repository.SingerRepository
}

func NewSingerService(repo repository.SingerRepository) *SingerService {
	return &SingerService{
		repo: repo,
	}
}

func (s *SingerService) CreateSinger(ctx context.Context, name string) (*model.Singer, error) {
	return s.repo.CreateSinger(ctx, name)
}
