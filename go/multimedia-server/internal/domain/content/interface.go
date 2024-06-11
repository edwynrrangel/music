package content

import (
	"context"
	"io"
)

type UseCase interface {
	Search(ctx context.Context, request *SearchRequest) (*SearchResponse, error)
	Get(ctx context.Context, id string) (*Content, error)
	Stream(ctx context.Context, id string) (io.ReadCloser, error)
}

type Repository interface {
	Get(ctx context.Context, id string) (*Content, error)
	Search(ctx context.Context, query string) ([]Content, error)
}
