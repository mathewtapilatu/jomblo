package response

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

func FromDomain(domain users.Domain) Users {
	return Users{
		Id:        domain.Id,
		Name:      domain.Name,
		Email:     domain.Email,
		Password:  domain.Password,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
