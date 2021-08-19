package chats

import (
	"context"
	"time"
)

type Domain struct {
	Id             int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
	ChannelChatsID int
	Message        string
}
type Usecase interface {
	ChannelChatsID(ctx context.Context, domain *Domain) (Domain, error)
	Message(ctx context.Context, domain *Domain) (Domain, error)
}

type Repository interface {
	RepoChatsID(ctx context.Context, domain *Domain) (Domain, error)
	RepoMessage(ctx context.Context, domain *Domain) (Domain, error)
}
