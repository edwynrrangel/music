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

func (r *PlaylistRequest) toPlayListRequest() *playlist.PlayListRequest {
	return &playlist.PlayListRequest{
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
