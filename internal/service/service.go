package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/openingwiki/wiki/internal/model"
	"github.com/openingwiki/wiki/internal/repository"
)

type Service struct { repos repository.Repos }

func New(repos repository.Repos) *Service { return &Service{repos: repos} }

func (s *Service) CreateAnime(title string) (*model.Anime, error) {
	title = strings.TrimSpace(title)
	if title == "" { return nil, errors.New("title is required") }
	return s.repos.Anime.CreateAnime(title)
}
func (s *Service) GetAnime(id int64) (*model.Anime, error) { return s.repos.Anime.GetAnime(id) }

func (s *Service) CreateSinger(name string) (*model.Singer, error) {
	name = strings.TrimSpace(name)
	if name == "" { return nil, errors.New("name is required") }
	return s.repos.Singer.CreateSinger(name)
}
func (s *Service) GetSinger(id int64) (*model.Singer, error) { return s.repos.Singer.GetSinger(id) }

func (s *Service) CreateOpening(animeID, singerID int64, t model.Type, title string, order int) (*model.Opening, error) {
	if _, err := s.repos.Anime.GetAnime(animeID); err != nil {
		if repository.IsNotFound(err) { return nil, fmt.Errorf("anime %d not found", animeID) }
		return nil, err
	}
	if _, err := s.repos.Singer.GetSinger(singerID); err != nil {
		if repository.IsNotFound(err) { return nil, fmt.Errorf("singer %d not found", singerID) }
		return nil, err
	}
	if title = strings.TrimSpace(title); title == "" { return nil, errors.New("title is required") }
	if t != model.TypeOpening && t != model.TypeEnding && t != model.TypeOST { return nil, errors.New("invalid type") }
	if order < 0 { return nil, errors.New("order_number must be >= 0") }
	return s.repos.Opening.CreateOpening(&model.Opening{AnimeID: animeID, SingerID: singerID, Type: t, Title: title, OrderNumber: order})
}
func (s *Service) GetOpening(id int64) (*model.Opening, error) { return s.repos.Opening.GetOpening(id) }
