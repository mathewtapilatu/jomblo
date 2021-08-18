package request

import (
	"jomblo/business/matches"
)

type Matches struct {
	UserID    int `json:"UserID"`
	MatchesID int `json:"MatchesID"`
}

func (req *Matches) ToDomain() *matches.Domain {
	return &matches.Domain{
		UserID:    req.UserID,
		MatchesID: req.MatchesID,
	}
}
