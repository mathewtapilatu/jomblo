package matches

import (
	"context"
	"jomblo/business/matches"

	"gorm.io/gorm"
)

type MysqlMatchesRepository struct {
	Conn *gorm.DB
}

func NewMysqlMatchesRepository(md *gorm.DB) matches.Repository {
	return &MysqlMatchesRepository{
		Conn: md,
	}
}

func (ur *MysqlMatchesRepository) RepoUserID(ctx context.Context, domain *matches.Domain) (matches.Domain, error) {
	rec := fromDomain(*domain)
	result := ur.Conn.Create(&rec)
	if result.Error != nil {
		return matches.Domain{}, result.Error
	}
	return rec.toDomain(), nil

}
func (ur *MysqlMatchesRepository) RepoMatchesID(ctx context.Context, domain *matches.Domain) (matches.Domain, error) {
	rec := fromDomain(matches.Domain{})
	return rec.toDomain(), nil
}
