package content

import (
	"context"
	"log"

	"github.com/edwynrrangel/music/go/content-server/internal/domain/content"
)

type adapter struct {
	UnimplementedContentServiceServer
	usecase content.UseCase
}

// NewAdapter function creates a new instance of adapter
func NewAdapter(usecase content.UseCase) *adapter {
	return &adapter{usecase: usecase}
}

// Search function returns a list of contents
func (a *adapter) Search(ctx context.Context, req *SearchRequest) (*SearchResponse, error) {
	log.Printf("Received request: %+v", req)
	paginated, err := a.usecase.Search(ctx, req.Query)
	if err != nil {
		return nil, err
	}

	return (*Paginated)(paginated).toResponse(), nil
}

func (a *adapter) Get(ctx context.Context, req *ContentRequest) (*Content, error) {
	log.Printf("Received request: %+v", req)
	got, err := a.usecase.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return (*ContentEntity)(got).toResponse(), nil
}
