package content

import (
	"context"
	"log"

	"github.com/edwynrrangel/grpc/go/multimedia-client/internal/adapter/grpc/content"
)

type usecase struct {
	content content.ContentServiceClient
}

// NewUseCase creates a new content usecase
func NewUseCase(content content.ContentServiceClient) UseCase {
	return &usecase{
		content: content,
	}
}

func (u *usecase) Search(ctx context.Context, data *SearchRequest) (*ContentList, error) {
	got, err := u.content.Search(ctx, (*content.SearchRequest)(data))
	if err != nil {
		return nil, err
	}
	return (*ContentList)(got), nil
}

func (u *usecase) Get(ctx context.Context, data *ContentRequest) (*Content, error) {
	got, err := u.content.Get(ctx, (*content.ContentRequest)(data))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return (*Content)(got), nil
}
