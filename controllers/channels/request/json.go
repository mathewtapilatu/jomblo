package request

import (
	"jomblo/business/channels"
)

type Channels struct {
	ChannelUserID    int `json:"ChannelUserID"`
	ChannelMatchesID int `json:"ChannelMatchesID"`
}

func (req *Channels) ToDomain() *channels.Domain {
	return &channels.Domain{
		ChannelUserID:    req.ChannelUserID,
		ChannelMatchesID: req.ChannelMatchesID,
	}
}
