package model

import "time"

type Singer struct {
	ID        int64
	Name      string
	CreatedAt time.Time
}
