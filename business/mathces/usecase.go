package mathces

import "context"

type MatchesUsecase struct {
	MatchesRepository Repository
}

func UserMatchesUsecase(Mr Repository) Usecase {
	return &MatchesUsecase{
		MatchesRepository: Mr,
	}
}

func (Mr MatchesUsecase) MatchesID(ctx context.Context, Id, UserID int) (Domain, error) {
	_, cancel := Mr.MatchesRepository.RepoMatchesID(ctx, Id, UserID)
	if cancel != nil {
		return Domain{}, nil
	}
	return Domain{}, nil
}

func (Mr MatchesUsecase) UserID(ctx context.Context, Id int) (Domain, error) {
	_, cancel := Mr.MatchesRepository.RepoUserID(ctx, Id)
	if cancel != nil {
		return Domain{}, nil
	}
	return Domain{}, nil
}
