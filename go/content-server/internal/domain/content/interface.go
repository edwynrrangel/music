package content

import (
	"context"

	"github.com/edwynrrangel/music/go/multimedia_server/internal/shared"
)

type UseCase interface {
	Search(ctx context.Context, query string) (*shared.Paginated[Content], error)
}

type Repository interface {
	Search(ctx context.Context, query string) ([]Content, error)
	Count(ctx context.Context, query string) (int64, error)
}
