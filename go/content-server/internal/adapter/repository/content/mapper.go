package content

import "github.com/edwynrrangel/music/go/multimedia_server/internal/domain/content"

func (l Locations) toEntity() content.Locations {
	var result content.Locations
	for _, item := range l {
		result = append(result, content.Location(item))
	}
	return result
}

func (c *Content) toEntity() *content.Content {
	return &content.Content{
		ID:        c.ID,
		Title:     c.Title,
		Artists:   c.Artists,
		Album:     c.Album,
		Genre:     c.Genre,
		Duration:  c.Duration,
		CoverURL:  c.CoverURL,
		Locations: (*Locations)(&c.Locations).toEntity(),
	}
}

func (c Contents) toArrayEntity() content.Contents {
	var result content.Contents
	for _, item := range c {
		result = append(result, *item.toEntity())
	}
	return result
}
