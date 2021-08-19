package request

import "jomblo/business/chats"

type Chats struct {
	ChannelChatsID int    `json:"ChannelChatsID"`
	Message        string `json:"Message"`
}

func (req *Chats) ToDomain() *chats.Domain {
	return &chats.Domain{
		ChannelChatsID: req.ChannelChatsID,
		Message:        req.Message,
	}
}
