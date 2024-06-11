package content

import "github.com/edwynrrangel/grpc/go/multimedia_server/internal/domain/content"

func (c *Content) toDomain() *content.Content {
	return &content.Content{
		ID:       c.ID,
		Title:    c.Title,
		Genre:    c.Genre,
		Creator:  c.Creator,
		Duration: c.Duration,
		Bucket:   c.Bucket,
	}
}

func toArrayDomain(contents []Content) []content.Content {
	var result []content.Content
	for _, c := range contents {
		result = append(result, *c.toDomain())
	}
	return result
}
