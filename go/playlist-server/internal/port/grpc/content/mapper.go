package content

import "github.com/edwynrrangel/grpc/go/playlist_server/internal/domain/content"

func convertToDomain(got *ContentResponse) *content.Content {
	return &content.Content{
		ID:       got.Id,
		Title:    got.Title,
		Genre:    got.Genre,
		Creator:  got.Creator,
		Duration: got.Duration,
	}
}
