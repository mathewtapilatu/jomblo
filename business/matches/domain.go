package matches

import (
	"context"
	"time"
)

type Domain struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	UserID    int
	MatchesID int
}

//matchesID = Userlike
//userID = user me
type Usecase interface {
	UserID(ctx context.Context, domain *Domain) (Domain, error)
	MatchesID(ctx context.Context, domain *Domain) (Domain, error)
}

type Repository interface {
	RepoUserID(ctx context.Context, domain *Domain) (Domain, error)
	RepoMatchesID(ctx context.Context, domain *Domain) (Domain, error)
}
