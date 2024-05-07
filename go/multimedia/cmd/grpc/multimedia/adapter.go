package multimedia

import (
	"context"
	"log"

	"github.com/edwynrrangel/grpc/go/multimedia/internal/multimedia"
	"github.com/edwynrrangel/grpc/go/multimedia/internal/proto"
)

type adapter struct {
	proto.UnimplementedMultimediaServiceServer
	usecase multimedia.UseCase
}

func NewAdapter(usecase multimedia.UseCase) *adapter {
	return &adapter{usecase: usecase}
}

func (a *adapter) SearchContent(ctx context.Context, req *proto.SearchRequest) (*proto.SearchResponse, error) {
	log.Printf("Received request: %+v", req)
	schema := multimedia.ConvertPBSearchRequestToSearchRequest(req)
	got, err := a.usecase.SearchContent(schema)
	if err != nil {
		return nil, err
	}

	return got.ToProto(), nil
}
