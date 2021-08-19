package chats

import "context"

type ChatsUsecase struct {
	ChatsRepository Repository
}

func UserChatsUsecase(Cc Repository) Usecase {
	return &ChatsUsecase{
		ChatsRepository: Cc,
	}
}

func (Cc ChatsUsecase) ChannelChatsID(ctx context.Context, domain *Domain) (Domain, error) {
	_, cancel := Cc.ChatsRepository.RepoChatsID(ctx, domain)
	if cancel != nil {
		return Domain{}, nil
	}
	return Domain{}, nil
}

func (Cc ChatsUsecase) Message(ctx context.Context, domain *Domain) (Domain, error) {
	_, cancel := Cc.ChatsRepository.RepoMessage(ctx, domain)
	if cancel != nil {
		return Domain{}, nil
	}
	return Domain{}, nil
}
