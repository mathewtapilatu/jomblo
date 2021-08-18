package matches

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
	MatchesID int
}

func (record *Matches) toDomain() matches.Domain {
	return matches.Domain{
		Id:        record.Id,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
		DeletedAt: record.DeletedAt,
		UserID:    record.UserID,
		MatchesID: record.MatchesID,
	}

}

func fromDomain(domain matches.Domain) *Matches {
	return &Matches{
		Id:        domain.Id,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		UserID:    domain.UserID,
		MatchesID: domain.MatchesID,
	}

}
