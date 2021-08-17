package request

import (
	"jomblo/business/matches"
	"time"
)

type Matches struct {
	UserID    string    `json:"UserID"`
	MathcesID string    `json:"MatchesID"`
	Name      string    `json:"Name"`
	CreatedAt time.Time `json:"Created_at"`
	UpdatedAt time.Time `json:"Updated_at"`
	DeletedAt time.Time `json:"Deleted_at"`
}

func (req *Matches) ToDomain() *matches.Domain {
	return &matches.Domain{
		UserID:    req.UserID,
		MatchesID: req.MathcesID,
		Name:      req.Name,
	}
}
