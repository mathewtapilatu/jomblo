package mathces

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
	MathcesID int
}

type Usecase interface {
	UserID(ctx context.Context, Id int) (Domain, error)
	MatchesID(ctx context.Context, Id, UserID int) (Domain, error)
}

type Repository interface {
	RepoUserID(ctx context.Context, Id int) (Domain, error)
	RepoMatchesID(ctx context.Context, Id, UserID int) (Domain, error)
}
