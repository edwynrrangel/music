package content

import (
	"context"
	"io"

	"github.com/edwynrrangel/grpc/go/multimedia_server/pkg/bucket"
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

func (u *usecase) Search(ctx context.Context, query string) (*SearchResponse, error) {
	contents, err := u.contentRepository.Search(ctx, query)
	if err != nil {
		return nil, err
	}

	return &SearchResponse{
		Data: contents,
	}, nil
}

func (u *usecase) Get(ctx context.Context, id string) (*Content, error) {
	content, err := u.contentRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (u *usecase) Stream(ctx context.Context, id string) (io.ReadCloser, error) {
	content, err := u.contentRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	file, err := u.bucketStrategy.DownloadFile(ctx, content.Bucket, id)
	if err != nil {
		return nil, err
	}

	return file, nil
}
