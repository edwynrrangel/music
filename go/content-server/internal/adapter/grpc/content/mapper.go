package content

import (
	"github.com/edwynrrangel/music/go/content-server/internal/domain/content"
	"github.com/edwynrrangel/music/go/content-server/internal/shared"
)

type (
	ContentEntity content.Content
	Paginated     shared.Paginated[content.Content]
	Locations     []content.Location
)

func (l *Locations) toResponse() []*Location {
	locations := make([]*Location, 0, len(*l))
	for _, item := range *l {
		locations = append(locations, &Location{
			Bucket: item.Bucket,
			Key:    item.Key,
		})
	}
	return locations
}

func (c *ContentEntity) toResponse() *Content {
	return &Content{
		Id:        c.ID,
		Title:     c.Title,
		Artists:   c.Artists,
		Album:     c.Album,
		Genre:     c.Genre,
		Duration:  c.Duration,
		CoverUrl:  c.CoverURL,
		Locations: (*Locations)(&c.Locations).toResponse(),
	}
}

func (p *Paginated) toResponse() *SearchResponse {
	contents := make([]*Content, 0, len(p.Data))
	for _, c := range p.Data {
		contents = append(contents, (*ContentEntity)(&c).toResponse())
	}

	return &SearchResponse{
		Total: p.Total,
		Data:  contents,
	}
}
