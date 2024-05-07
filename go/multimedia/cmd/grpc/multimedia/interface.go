package multimedia

import (
	"context"

	"github.com/edwynrrangel/grpc/go/multimedia/internal/proto"
)

type Adapter interface {
	SearchContent(ctx context.Context, req *proto.SearchRequest) (*proto.SearchResponse, error)
}
