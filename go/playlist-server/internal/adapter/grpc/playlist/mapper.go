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
		UserID: c.UserId,
		Name:   c.Name,
	}
}

func (e *PlaylistEntity) toResponse() *Playlist {
	return &Playlist{
		Id:   e.ID,
		Name: e.Name,
	}
}

func (l *ListRequest) toRequest() *playlist.ListRequest {
	return &playlist.ListRequest{
		UserID: l.UserId,
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

/* func (c *Content) toContentRequest() *playlist.ContentRequest {
	return &playlist.ContentRequest{
		ID:       c.Id,
		Title:    c.Title,
		Genre:    c.Genre,
		Creator:  c.Creator,
		Duration: c.Duration,
	}
}

func (r *PlaylistRequest) toPlayListRequest() *playlist.PlaylistRequest {
	contents := make([]playlist.ContentRequest, 0)
	for _, content := range r.Contents {
		contents = append(contents, *content.toContentRequest())
	}
	return &playlist.PlaylistRequest{
		ID:        r.Id,
		UserID:    r.UserID,
		Name:      r.Name,
		Contents:  contents,
		Operation: playlist.OperationKey(r.Operation),
	}
}

func convertToContentResponse(content *playlist.Content) *Content {
	return &Content{
		Id:       content.ID,
		Title:    content.Title,
		Genre:    content.Genre,
		Creator:  content.Creator,
		Duration: content.Duration,
	}
}

func convertToPlaylistResponse(playlist *playlist.Playlist) *PlaylistResponse {
	contents := make([]*Content, 0)
	for _, content := range playlist.Contents {
		contents = append(contents, convertToContentResponse(&content))
	}
	return &PlaylistResponse{
		Id:       playlist.ID,
		Name:     playlist.Name,
		Contents: contents,
	}
}

func convertToListPlaylistResponse_Playlist(playlist *playlist.Playlist) *ListPlaylistResponse_Playlist {
	return &ListPlaylistResponse_Playlist{
		Id:   playlist.ID,
		Name: playlist.Name,
	}
}

func convertToListPlaylistResponse(list []playlist.Playlist) *ListPlaylistResponse {
	listPB := make([]*ListPlaylistResponse_Playlist, 0)
	for _, item := range list {
		listPB = append(listPB, convertToListPlaylistResponse_Playlist(&item))
	}
	return &ListPlaylistResponse{
		Data: listPB,
	}
}
*/
