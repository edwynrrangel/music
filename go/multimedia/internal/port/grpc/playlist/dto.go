package playlist

import "github.com/edwynrrangel/grpc/go/multimedia/internal/domain/playlist"

func (r *PlaylistRequest) toPlayListRequest() *playlist.PlayListRequest {
	return &playlist.PlayListRequest{
		ID:        r.Id,
		UserID:    r.UserID,
		ContentID: r.ContentID,
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

func convertToPlayListResponse(playlist *playlist.PlayList) *PlaylistResponse {
	contents := make([]*Content, 0)
	for _, content := range playlist.Content {
		contents = append(contents, convertToContentResponse(&content))
	}
	return &PlaylistResponse{
		UserID: playlist.UserID,
		Name:   playlist.Name,
		Data:   contents,
	}
}
