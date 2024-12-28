package content

import "github.com/edwynrrangel/music/go/multimedia_server/internal/domain/content"

func (c *Content) toEntity() *content.Content {
	return &content.Content{
		ID:       c.ID,
		Title:    c.Title,
		Artists:  c.Artists,
		Album:    c.Album,
		Genre:    c.Genre,
		Duration: c.Duration,
		CoverURL: c.CoverURL,
	}
}

func (c Contents) toArrayEntity() []content.Content {
	var result []content.Content
	for _, item := range c {
		result = append(result, *item.toEntity())
	}
	return result
}
