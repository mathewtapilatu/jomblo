package response

import (
	"jomblo/business/matches"
	"time"
)

type Matches struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	UserID    int
	MathcesID int
}

func FromDomain(domain matches.Domain) Matches {
	return Matches{
		Id:        domain.Id,
		UserID:    domain.UserID,
		MathcesID: domain.MatchesID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
