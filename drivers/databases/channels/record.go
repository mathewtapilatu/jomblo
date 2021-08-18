package channels

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

func (record *Channels) toDomain() channels.Domain {
	return channels.Domain{
		Id:               record.Id,
		CreatedAt:        record.CreatedAt,
		UpdatedAt:        record.UpdatedAt,
		DeletedAt:        record.DeletedAt,
		ChannelUserID:    record.ChannelUserID,
		ChannelMatchesID: record.ChannelMatchesID,
	}

}

func fromDomain(domain channels.Domain) *Channels {
	return &Channels{
		Id:               domain.Id,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
		DeletedAt:        domain.DeletedAt,
		ChannelUserID:    domain.ChannelUserID,
		ChannelMatchesID: domain.ChannelMatchesID,
	}

}
