package playlist

import "github.com/edwynrrangel/grpc/go/multimedia/internal/domain/playlist"

type (
	PlaylistDTO struct {
		ID   string
		Name string
	}
	ListPlaylistsDTO struct {
		UserID string
		Data   []PlaylistDTO
	}
)

func (r *PlaylistRequest) toPlayListRequest() *playlist.PlaylistRequest {
	return &playlist.PlaylistRequest{
		ID:        r.Id,
		UserID:    r.UserID,
		ContentID: r.ContentID,
		Name:      r.Name,
	}
}

func convertToContentResponse(content *playlist.Content) *Content {
	return &Content{
		Id:      content.ID,
		Title:   content.Title,
		Genre:   content.Genre,
		Creator: content.Creator,
	}
}

func convertToPlayListResponse(playlist *playlist.Playlist) *PlaylistResponse {
	contents := make([]*Content, 0)
	for _, content := range playlist.Content {
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

func (r *RemovePlaylistRequest) toRemovePlaylistRequest() *playlist.RemovePlaylistRequest {
	return &playlist.RemovePlaylistRequest{
		ID:     r.Id,
		UserID: r.UserID,
	}
}
