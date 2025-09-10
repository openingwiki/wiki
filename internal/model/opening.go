package model

import "time"

type OpeningType string

const (
	TypeOpening OpeningType = "opening"
	TypeEnding  OpeningType = "ending"
)

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
