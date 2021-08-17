package users

import (
	"jomblo/business/users"
	"time"
)

type Users struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Name      string
	Email     string
	Alamat    string
	Password  string
	Umur      int
}

func (record *Users) toDomain() users.Domain {
	return users.Domain{
		Id:        record.Id,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
		DeletedAt: record.DeletedAt,
		Name:      record.Name,
		Email:     record.Email,
		Alamat:    record.Alamat,
		Password:  record.Password,
		Umur:      record.Umur,
	}

}

func fromDomain(domain users.Domain) *Users {
	return &Users{
		Id:        domain.Id,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Name:      domain.Name,
		Email:     domain.Email,
		Alamat:    domain.Alamat,
		Password:  domain.Password,
		Umur:      domain.Umur,
	}

}
