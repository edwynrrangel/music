package content

import (
	"context"
	"io"
)

type UseCase interface {
	Get(ctx context.Context, id string) (*Content, error)
	Search(ctx context.Context, query string) (*SearchResponse, error)
	Stream(ctx context.Context, id string) (io.ReadCloser, error)
}

type Repository interface {
	Get(ctx context.Context, id string) (*Content, error)
	Search(ctx context.Context, query string) ([]Content, error)
}
