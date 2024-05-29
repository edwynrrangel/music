package content

import (
	"context"
	"io"
	"log"

	"github.com/edwynrrangel/grpc/go/multimedia/internal/content"
	"github.com/edwynrrangel/grpc/go/multimedia/internal/proto"
)

type adapter struct {
	proto.UnimplementedMultimediaServiceServer
	usecase content.UseCase
}

func NewAdapter(usecase content.UseCase) *adapter {
	return &adapter{usecase: usecase}
}

func (a *adapter) SearchContent(ctx context.Context, req *proto.SearchRequest) (*proto.SearchResponse, error) {
	log.Printf("Received request: %+v", req)
	schema := content.ConvertPBSearchRequestToSearchRequest(req)
	got, err := a.usecase.SearchContent(ctx, schema)
	if err != nil {
		return nil, err
	}

	return got.ToProto(), nil
}

func (a *adapter) GetContent(ctx context.Context, req *proto.StreamRequest) (*proto.ContentResponse, error) {
	log.Printf("Received request: %+v", req)
	got, err := a.usecase.GetContent(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return got.ToProto(), nil
}

func (a *adapter) StreamContent(req *proto.StreamRequest, stream proto.MultimediaService_StreamContentServer) error {
	log.Printf("Received request: %+v", req)

	file, err := a.usecase.StreamContent(stream.Context(), req.Id)
	if err != nil {
		return err
	}

	defer file.Close()

	buffer := make([]byte, 64*1024)
	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		if err := stream.Send(&proto.StreamResponse{Data: buffer[:n]}); err != nil {
			return err
		}
	}

	return nil
}
