package channels

import (
	"context"
	"jomblo/business/channels"

	"gorm.io/gorm"
)

type MysqlChannelsRepository struct {
	Conn *gorm.DB
}

func NewMysqlChannelsRepository(cc *gorm.DB) channels.Repository {
	return &MysqlChannelsRepository{
		Conn: cc,
	}
}

func (cr *MysqlChannelsRepository) RepoChannels(ctx context.Context, domain *channels.Domain) (channels.Domain, error) {
	rec := fromDomain(*domain)
	result := cr.Conn.Create(&rec)
	if result.Error != nil {
		return channels.Domain{}, result.Error
	}
	return rec.toDomain(), nil

}
func (cr *MysqlChannelsRepository) RepoChannelsID(ctx context.Context, domain *channels.Domain) (channels.Domain, error) {
	rec := fromDomain(channels.Domain{})
	return rec.toDomain(), nil
}
