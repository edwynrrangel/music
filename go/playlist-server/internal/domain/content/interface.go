package content

import "context"

type Adapter interface {
	Get(ctx context.Context, id string) (*Content, error)
}
