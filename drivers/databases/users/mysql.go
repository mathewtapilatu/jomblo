package users

import (
	"context"
	"jomblo/business/users"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(db *gorm.DB) users.Repository {
	return &MysqlUserRepository{
		Conn: db,
	}
}

func (ur *MysqlUserRepository) UserRegis(ctx context.Context, domain *users.Domain) (users.Domain, error) {
	rec := fromDomain(*domain)
	result := ur.Conn.Create(&rec)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}
	return rec.toDomain(), nil

}
func (ur *MysqlUserRepository) UserLogin(ctx context.Context, email, password string) (users.Domain, error) {
	rec := fromDomain(users.Domain{})
	return rec.toDomain(), nil
}
