package multimedia

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

func (u *usecase) SearchContent(ctx context.Context, request *SearchRequest) (*SearchResponse, error) {
	contents, err := u.repo.SearchContent(ctx, request.Query)
	if err != nil {
		return nil, err
	}

	return &SearchResponse{
		Data: contents,
	}, nil
}

func (u *usecase) StreamContent(ctx context.Context, id string) (io.ReadCloser, error) {
	file, err := u.bucketStrategy.DownloadFile(ctx, "multimedia", id)
	if err != nil {
		return nil, err
	}

	return file, nil
}
