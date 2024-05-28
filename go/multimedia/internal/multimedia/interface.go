package multimedia

import (
	"context"
	"io"
)

type UseCase interface {
	SearchContent(ctx context.Context, request *SearchRequest) (*SearchResponse, error)
	StreamContent(ctx context.Context, id string) (io.ReadCloser, error)
}

type Repository interface {
	SearchContent(ctx context.Context, query string) ([]Content, error)
	StreamContent(ctx context.Context, id string) (Content, error)
}
