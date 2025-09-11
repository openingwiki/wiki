package model

import "time"

type Opening struct {
	ID          int64
	AnimeId     int64
	SingerId    int64
	Type        OpeningType
	Title       string
	OrderNumber int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
