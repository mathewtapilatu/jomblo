package channels

import "context"

type ChannelsUsecase struct {
	ChannelsRepository Repository
}

func UserChannelsUsecase(Cr Repository) Usecase {
	return &ChannelsUsecase{
		ChannelsRepository: Cr,
	}
}

func (Cr ChannelsUsecase) ChannelsUserID(ctx context.Context, domain *Domain) (Domain, error) {
	_, cancel := Cr.ChannelsRepository.RepoChannels(ctx, domain)
	if cancel != nil {
		return Domain{}, nil
	}
	return Domain{}, nil
}

func (Cr ChannelsUsecase) ChannelsMatchesID(ctx context.Context, domain *Domain) (Domain, error) {
	_, cancel := Cr.ChannelsRepository.RepoChannelsID(ctx, domain)
	if cancel != nil {
		return Domain{}, nil
	}
	return Domain{}, nil
}
