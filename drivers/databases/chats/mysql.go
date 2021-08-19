package chats

import (
	"context"
	"jomblo/business/chats"

	"gorm.io/gorm"
)

type MysqlChatsRepository struct {
	Conn *gorm.DB
}

func NewMysqlChatsRepository(cm *gorm.DB) chats.Repository {
	return &MysqlChatsRepository{
		Conn: cm,
	}
}

func (cm *MysqlChatsRepository) RepoChatsID(ctx context.Context, domain *chats.Domain) (chats.Domain, error) {
	rec := fromDomain(*domain)
	result := cm.Conn.Create(&rec)
	if result.Error != nil {
		return chats.Domain{}, result.Error
	}
	return rec.toDomain(), nil

}
func (cm *MysqlChatsRepository) RepoMessage(ctx context.Context, domain *chats.Domain) (chats.Domain, error) {
	rec := fromDomain(chats.Domain{})
	return rec.toDomain(), nil
}
