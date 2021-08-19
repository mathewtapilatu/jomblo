package chats

import (
	"jomblo/business/chats"
	"time"
)

type Chats struct {
	Id             int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
	ChannelChatsID int
	Message        string
}

func (record *Chats) toDomain() chats.Domain {
	return chats.Domain{
		Id:             record.Id,
		CreatedAt:      record.CreatedAt,
		UpdatedAt:      record.UpdatedAt,
		DeletedAt:      record.DeletedAt,
		ChannelChatsID: record.ChannelChatsID,
		Message:        record.Message,
	}

}

func fromDomain(domain chats.Domain) *Chats {
	return &Chats{
		Id:             domain.Id,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
		DeletedAt:      domain.DeletedAt,
		ChannelChatsID: domain.ChannelChatsID,
		Message:        domain.Message,
	}

}
