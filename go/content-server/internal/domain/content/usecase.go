package content

import (
	"context"

	"github.com/edwynrrangel/music/go/content-server/internal/shared"
)

type usecase struct {
	contentRepository Repository
}

func NewUseCase(contentRepository Repository) UseCase {
	return &usecase{
		contentRepository: contentRepository,
	}
}

func (u *usecase) Search(ctx context.Context, query string) (*shared.Paginated[Content], error) {
	total, err := u.contentRepository.Count(ctx, query)
	if err != nil {
		return nil, err
	}

	contents, err := u.contentRepository.Search(ctx, query)
	if err != nil {
		return nil, err
	}

	return &shared.Paginated[Content]{
		Total: total,
		Data:  contents,
	}, nil
}

func (u *usecase) Get(ctx context.Context, id string) (*Content, error) {
	content, err := u.contentRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return content, nil
}
