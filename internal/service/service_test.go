package service

import (
	"testing"

	"github.com/openingwiki/wiki/internal/repository"
)

func TestCreateAnimeAndFetch(t *testing.T) {
	repo := repository.NewInMemory()
	svc := New(repo)

	created, err := svc.CreateAnime("Naruto")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got, err := svc.GetAnime(created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Title != "Naruto" {
		t.Fatalf("want title Naruto, got %s", got.Title)
	}
}

func TestCreateOpeningValidation(t *testing.T) {
	repo := repository.NewInMemory()
	svc := New(repo)

	if _, err := svc.CreateOpening(1, 1, "opening", "Song", 1); err == nil {
		t.Fatalf("expected error due to missing anime/singer")
	}
}
