package playlist

import (
	"context"
)

type usecase struct {
	playListRepository Repository
}

func NewUseCase(playListRepository Repository) UseCase {
	return &usecase{
		playListRepository: playListRepository,
	}
}

func (u *usecase) Get(ctx context.Context, request PlayListRequest) (*PlayList, error) {
	return nil, nil
}

func (u *usecase) Add(ctx context.Context, request PlayListRequest) error {
	return nil
}

func (u *usecase) Remove(ctx context.Context, request PlayListRequest) error {
	return nil
}
