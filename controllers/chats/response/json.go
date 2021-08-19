package response

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

func FromDomain(domain chats.Domain) Chats {
	return Chats{
		Id:             domain.Id,
		ChannelChatsID: domain.ChannelChatsID,
		Message:        domain.Message,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
	}
}
