package multimedia

import "context"

type usecase struct {
	repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &usecase{repo: repo}
}

func (u *usecase) SearchContent(ctx context.Context, request *SearchRequest) (*SearchResponse, error) {
	contents, err := u.repo.SearchContent(ctx, request.Query)
	if err != nil {
		return nil, err
	}

	return &SearchResponse{
		Data: contents,
	}, nil
}
