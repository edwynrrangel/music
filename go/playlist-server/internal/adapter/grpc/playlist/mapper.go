package playlist

import (
	"github.com/edwynrrangel/grpc/go/playlist_server/internal/domain/playlist"
	"github.com/edwynrrangel/grpc/go/playlist_server/internal/shared"
)

type (
	PlaylistEntity playlist.Playlist
	Paginated      shared.Paginated[playlist.Playlist]
)

func (c *CreateRequest) toRequest() *playlist.CreateRequest {
	return &playlist.CreateRequest{
		UserId: c.UserId,
		Name:   c.Name,
	}
}

func (e *PlaylistEntity) toResponse() *Playlist {
	return &Playlist{
		Id:      e.ID,
		Name:    e.Name,
		Content: e.Data,
	}
}

func (l *ListRequest) toRequest() *playlist.ListRequest {
	return &playlist.ListRequest{
		UserId: l.UserId,
		Query:  l.Query,
		Page:   l.Page,
		Limit:  l.Limit,
	}
}

func (p *Paginated) toResponse() *ListResponse {
	if p == nil {
		return nil
	}
	playlists := make([]*Playlist, 0, len(p.Data))
	for _, item := range p.Data {
		playlists = append(playlists, (*PlaylistEntity)(&item).toResponse())
	}

	return &ListResponse{
		Page:  p.Page,
		Limit: p.Limit,
		Total: p.Total,
		Data:  playlists,
	}
}
