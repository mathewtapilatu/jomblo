package response

import (
	"jomblo/business/channels"
	"time"
)

type Channels struct {
	Id               int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
	ChannelUserID    int
	ChannelMatchesID int
}

func FromDomain(domain channels.Domain) Channels {
	return Channels{
		Id:               domain.Id,
		ChannelUserID:    domain.ChannelUserID,
		ChannelMatchesID: domain.ChannelMatchesID,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
	}
}
