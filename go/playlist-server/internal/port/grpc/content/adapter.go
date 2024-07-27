package content

import (
	context "context"

	"github.com/edwynrrangel/grpc/go/playlist_server/internal/domain/content"
)

type adapter struct {
	content ContentServiceClient
}

// NewAdapter ...
func NewAdapter(content ContentServiceClient) content.Adapter {
	return &adapter{content: content}
}

// Get ...
func (a *adapter) Get(ctx context.Context, id string) (*content.Content, error) {
	got, err := a.content.Get(ctx, &ContentRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return convertToDomain(got), nil
}
