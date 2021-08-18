package matches

import "context"

type MatchesUsecase struct {
	MatchesRepository Repository
}

func UserMatchesUsecase(Mr Repository) Usecase {
	return &MatchesUsecase{
		MatchesRepository: Mr,
	}
}

func (Mr MatchesUsecase) MatchesID(ctx context.Context, domain *Domain) (Domain, error) {
	_, cancel := Mr.MatchesRepository.RepoMatchesID(ctx, domain)
	if cancel != nil {
		return Domain{}, nil
	}
	return Domain{}, nil
}

func (Mr MatchesUsecase) UserID(ctx context.Context, domain *Domain) (Domain, error) {
	_, cancel := Mr.MatchesRepository.RepoUserID(ctx, domain)
	if cancel != nil {
		return Domain{}, nil
	}
	return Domain{}, nil
}
