package service

import (
	"context"
	"errors"

	"github.com/openingwiki/wiki/internal/model"
	"github.com/openingwiki/wiki/internal/repository"
)

var ErrOpeningExist = errors.New("opening already exist")

type OpeningService struct {
	repo repository.OpeningRepository
}

func NewOpeningService(repo repository.OpeningRepository) *OpeningService {
	return &OpeningService{repo: repo}
}
func (s *OpeningService) CreateOpening(ctx context.Context,
	animeID int64,
	singerID int64,
	openingType model.OpeningType,
	title string,
	orderNumber int64) (*model.Opening, error) {
	return s.repo.CreateOpening(ctx, animeID, singerID, openingType, title, orderNumber)
}
