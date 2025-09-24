package formatter

import (
	"time"

	"github.com/openingwiki/wiki/internal/model"
)

type CreateSingerRequest struct {
	Name string `json:"name" binding:"required"` 
}

type SingerResponse struct {
	ID	      int64		`json:"id"`
	Name	  string	`json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func NewSingerResponseFromDomain(m *model.Singer) *SingerResponse {
	return &SingerResponse{
		ID:		   m.ID,
		Name:	   m.Name,
		CreatedAt: m.CreatedAt,
	}
}
