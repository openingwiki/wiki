package model

import "time"

// Anime represents an anime title in the database
type Anime struct {
	ID        int64
	Title     string
	CreatedAt time.Time
}
