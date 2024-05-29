package content

import (
	"context"
	"io"

	"github.com/edwynrrangel/grpc/go/multimedia/pkg/bucket"
)

type usecase struct {
	bucketStrategy bucket.Strategy
	repo           Repository
}

func NewUseCase(repo Repository, bucketStrategy bucket.Strategy) UseCase {
	return &usecase{
		bucketStrategy: bucketStrategy,
		repo:           repo,
	}
}

func (u *usecase) Search(ctx context.Context, request *SearchRequest) (*SearchResponse, error) {
	contents, err := u.repo.Search(ctx, request.Query)
	if err != nil {
		return nil, err
	}

	return &SearchResponse{
		Data: contents,
	}, nil
}

func (u *usecase) Get(ctx context.Context, id string) (*Content, error) {
	content, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &content, nil
}

func (u *usecase) Stream(ctx context.Context, id string) (io.ReadCloser, error) {
	content, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	file, err := u.bucketStrategy.DownloadFile(ctx, content.Bucket, id)
	if err != nil {
		return nil, err
	}

	return file, nil
}
