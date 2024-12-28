package content

import (
	"context"

	"github.com/edwynrrangel/music/go/multimedia_server/internal/shared"
	"github.com/edwynrrangel/music/go/multimedia_server/pkg/bucket"
)

type usecase struct {
	bucketStrategy    bucket.Strategy
	contentRepository Repository
}

func NewUseCase(bucketStrategy bucket.Strategy, contentRepository Repository) UseCase {
	return &usecase{
		bucketStrategy:    bucketStrategy,
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
