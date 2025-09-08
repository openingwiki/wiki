package repository

import (
	"errors"
	"sync"
	"time"

	"github.com/openingwiki/wiki/internal/model"
)

var (
	errNotFound = errors.New("not found")
)

type InMemory struct {
	mu          sync.RWMutex
	animeByID   map[int64]*model.Anime
	singerByID  map[int64]*model.Singer
	openingByID map[int64]*model.Opening
	nextID      int64
}

func NewInMemory() *InMemory {
	return &InMemory{
		animeByID:   make(map[int64]*model.Anime),
		singerByID:  make(map[int64]*model.Singer),
		openingByID: make(map[int64]*model.Opening),
		nextID:      1,
	}
}

func (r *InMemory) next() int64 { r.nextID++; return r.nextID - 1 }

func (r *InMemory) CreateAnime(title string) (*model.Anime, error) {
	r.mu.Lock(); defer r.mu.Unlock()
	id := r.next()
	a := &model.Anime{ID: id, Title: title, CreatedAt: time.Now()}
	r.animeByID[id] = a
	return a, nil
}
func (r *InMemory) GetAnime(id int64) (*model.Anime, error) {
	r.mu.RLock(); defer r.mu.RUnlock()
	a, ok := r.animeByID[id]
	if !ok { return nil, errNotFound }
	return a, nil
}

func (r *InMemory) CreateSinger(name string) (*model.Singer, error) {
	r.mu.Lock(); defer r.mu.Unlock()
	id := r.next()
	s := &model.Singer{ID: id, Name: name, CreatedAt: time.Now()}
	r.singerByID[id] = s
	return s, nil
}
func (r *InMemory) GetSinger(id int64) (*model.Singer, error) {
	r.mu.RLock(); defer r.mu.RUnlock()
	s, ok := r.singerByID[id]
	if !ok { return nil, errNotFound }
	return s, nil
}

func (r *InMemory) CreateOpening(o *model.Opening) (*model.Opening, error) {
	r.mu.Lock(); defer r.mu.Unlock()
	o.ID = r.next()
	o.CreatedAt = time.Now()
	r.openingByID[o.ID] = o
	return o, nil
}
func (r *InMemory) GetOpening(id int64) (*model.Opening, error) {
	r.mu.RLock(); defer r.mu.RUnlock()
	o, ok := r.openingByID[id]
	if !ok { return nil, errNotFound }
	return o, nil
}

func IsNotFound(err error) bool { return errors.Is(err, errNotFound) }
