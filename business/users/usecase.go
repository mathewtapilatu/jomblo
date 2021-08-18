package users

import (
	"context"
	"log"
)

type UserUsecase struct {
	UserRepository Repository
}

func NewUserUseCase(Ur Repository) Usecase {
	return &UserUsecase{
		UserRepository: Ur,
	}
}

func (Uc UserUsecase) Login(ctx context.Context, email, password string) (Domain, error) {
	userlogin, err := Uc.UserRepository.UserLogin(ctx, email, password)
	if err != nil {
		return Domain{}, err
	}
	return userlogin, nil
}
func (Uc UserUsecase) Regis(ctx context.Context, domain *Domain) (Domain, error) {

	err := domain.valid()
	if err != nil {
		return Domain{}, err
	}

	userregister, err := Uc.UserRepository.UserRegis(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	log.Println(domain)

	return userregister, nil

}
