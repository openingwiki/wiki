package repository

import "github.com/openingwiki/wiki/internal/model"

// Split repositories per aggregate
type AnimeRepository interface {
    CreateAnime(title string) (*model.Anime, error)
    GetAnime(id int64) (*model.Anime, error)
}

type SingerRepository interface {
    CreateSinger(name string) (*model.Singer, error)
    GetSinger(id int64) (*model.Singer, error)
}

type OpeningRepository interface {
    CreateOpening(o *model.Opening) (*model.Opening, error)
    GetOpening(id int64) (*model.Opening, error)
}

// Repos bundles all repositories for wiring
type Repos struct {
    Anime   AnimeRepository
    Singer  SingerRepository
    Opening OpeningRepository
}


