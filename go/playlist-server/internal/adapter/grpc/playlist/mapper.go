package playlist

import (
	"github.com/edwynrrangel/grpc/go/playlist_server/internal/domain/playlist"
	"github.com/edwynrrangel/grpc/go/playlist_server/internal/shared"
)

type (
	ContentEntity  playlist.Content
	PlaylistEntity playlist.Playlist
	Paginated      shared.Paginated[playlist.Playlist]
)

func (c *CreateRequest) toRequest() *playlist.CreateRequest {
	return &playlist.CreateRequest{
		UserId: c.UserId,
		Name:   c.Name,
		Mode:   c.Mode,
	}
}

func (c *ContentEntity) toResponse() *Content {
	return &Content{
		Id:    c.ID,
		Order: c.Order,
	}
}

func (e *PlaylistEntity) toResponse() *Playlist {
	if e == nil {
		return nil
	}

	contents := make([]*Content, 0, len(e.Data))
	for _, item := range e.Data {
		contents = append(contents, (*ContentEntity)(&item).toResponse())
	}

	return &Playlist{
		Id:       e.ID,
		Name:     e.Name,
		Contents: contents,
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

func (p *PartyModeRequest) toRequest() *playlist.PartyModeRequest {
	var action playlist.Action
	switch x := p.GetAction().(type) {
	case *PartyModeRequest_AddContent:
		action.AddContent = &playlist.AddContentEvent{
			ContentId: x.AddContent.GetContentId(),
			Order:     x.AddContent.GetOrder(),
		}
	case *PartyModeRequest_SkipContent:
		action.SkipContent = &playlist.SkipContentEvent{
			ForceSkip: x.SkipContent.GetForceSkip(),
		}
	}

	return &playlist.PartyModeRequest{
		PlaylistId: p.GetPlaylistId(),
		Action:     action,
	}
}
