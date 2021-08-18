package users

import (
	"context"
	"errors"
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

func (d Domain) valid() (err error) {

	if d.Name == "" {
		return errors.New("gak isok")
	}
	return nil
}

type Usecase interface {
	Login(ctx context.Context, email, password string) (Domain, error)
	Regis(ctx context.Context, domain *Domain) (Domain, error)
}

type Repository interface {
	UserRegis(ctx context.Context, domain *Domain) (Domain, error)
	UserLogin(ctx context.Context, email, password string) (Domain, error)
	GetByParam(ctx context.Context, domain *Domain) (Domain, error)
}
