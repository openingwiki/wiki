package model

// OpeningType represents the type of anime theme
// @enum: opening, ending, ost
type OpeningType string

const (
	TypeOpening OpeningType = "opening"
	TypeEnding  OpeningType = "ending"
	TypeOst     OpeningType = "ost"
)
