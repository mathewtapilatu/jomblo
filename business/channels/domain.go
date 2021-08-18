package channels

import (
	"context"
	"time"
)

type Domain struct {
	Id               int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
	ChannelUserID    int
	ChannelMatchesID int
}
type Usecase interface {
	ChannelsUserID(ctx context.Context, domain *Domain) (Domain, error)
	ChannelsMatchesID(ctx context.Context, domain *Domain) (Domain, error)
}

type Repository interface {
	RepoChannels(ctx context.Context, domain *Domain) (Domain, error)
	RepoChannelsID(ctx context.Context, domain *Domain) (Domain, error)
}
