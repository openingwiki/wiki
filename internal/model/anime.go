package model

import "time"

type Anime struct {
	ID        int64
	Title     string
	CreatedAt time.Time
}
