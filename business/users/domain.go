package users

import (
	"context"
	"time"
)

type Domain struct {
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

type Usecase interface {
	Login(ctx context.Context, email, password string) (Domain, error)
	Regis(ctx context.Context, domain *Domain) (Domain, error)
}

type Repository interface {
	UserRegis(ctx context.Context, domain *Domain) (Domain, error)
	UserLogin(ctx context.Context, email, password string) (Domain, error)
}
