package playlist

import "github.com/edwynrrangel/grpc/go/multimedia_client/internal/adapter/grpc/playlist"

func convertToDataList(got *playlist.ListPlaylistResponse_Playlist) DataList {
	return DataList{
		ID:   got.Id,
		Name: got.Name,
	}
}

func convertToListResponse(got *playlist.ListPlaylistResponse) *ListResponse {
	var result []DataList
	for _, item := range got.Data {
		result = append(result, convertToDataList(item))
	}
	return &ListResponse{
		Data: result,
	}
}
