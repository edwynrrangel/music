package content

import (
	"context"
	"io"
	"log"

	"github.com/edwynrrangel/grpc/go/multimedia_server/internal/domain/content"
)

type adapter struct {
	UnimplementedContentServiceServer
	usecase content.UseCase
}

// NewAdapter function creates a new instance of adapter
func NewAdapter(usecase content.UseCase) *adapter {
	return &adapter{usecase: usecase}
}

// Get function returns a content
func (a *adapter) Get(ctx context.Context, req *ContentRequest) (*ContentResponse, error) {
	log.Printf("Received request: %+v", req)
	got, err := a.usecase.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return convertToContent(got), nil
}

// Search function returns a list of contents
func (a *adapter) Search(ctx context.Context, req *SearchRequest) (*SearchResponse, error) {
	log.Printf("Received request: %+v", req)
	got, err := a.usecase.Search(ctx, req.Query)
	if err != nil {
		return nil, err
	}

	return convertToSearchResponse(got), nil
}

// Stream function streams a content
func (a *adapter) Stream(req *ContentRequest, stream ContentService_StreamServer) error {
	log.Printf("Received request: %+v", req)
	file, err := a.usecase.Stream(stream.Context(), req.Id)
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := make([]byte, 64*1024)
	cont := 0
	for {

		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		if err := stream.Send(&StreamResponse{Data: buffer[:n]}); err != nil {
			return err
		}
		cont++
		log.Printf("chunk %d sent", cont)
	}

	return nil
}
