package content

import (
	"github.com/edwynrrangel/music/go/multimedia_server/internal/domain/content"
	"github.com/edwynrrangel/music/go/multimedia_server/internal/shared"
)

type ContentEntity content.Content
type Paginated shared.Paginated[content.Content]

func (c *ContentEntity) toResponse() *Content {
	return &Content{
		Id:       c.ID,
		Title:    c.Title,
		Artists:  c.Artists,
		Album:    c.Album,
		Genre:    c.Genre,
		Duration: c.Duration,
		CoverUrl: c.CoverURL,
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
