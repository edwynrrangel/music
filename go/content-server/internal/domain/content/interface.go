package content

import (
	"context"

	"github.com/edwynrrangel/music/go/content-server/internal/shared"
)

type UseCase interface {
	Search(ctx context.Context, query string) (*shared.Paginated[Content], error)
	Get(ctx context.Context, id string) (*Content, error)
}

type Repository interface {
	Search(ctx context.Context, query string) (Contents, error)
	Count(ctx context.Context, query string) (int64, error)
	Get(ctx context.Context, id string) (*Content, error)
}
