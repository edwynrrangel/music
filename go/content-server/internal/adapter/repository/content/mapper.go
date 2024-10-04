package content

import "github.com/edwynrrangel/grpc/go/multimedia_server/internal/domain/content"

func (c *Content) toEntity() *content.Content {
	return &content.Content{
		ID:       c.ID,
		Title:    c.Title,
		Genre:    c.Genre,
		Creator:  c.Creator,
		Duration: c.Duration,
		Bucket:   c.Bucket,
	}
}

func toArrayEntity(contents []Content) []content.Content {
	var result []content.Content
	for _, c := range contents {
		result = append(result, *c.toEntity())
	}
	return result
}
