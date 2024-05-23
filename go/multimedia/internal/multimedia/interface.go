package multimedia

import "context"

type UseCase interface {
	SearchContent(ctx context.Context, request *SearchRequest) (*SearchResponse, error)
}

type Repository interface {
	SearchContent(ctx context.Context, query string) ([]Content, error)
}
