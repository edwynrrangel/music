package playlist

import "github.com/edwynrrangel/grpc/go/multimedia-client/internal/adapter/grpc/playlist"

type (
	CreateRequest   playlist.CreateRequest
	ListRequest     playlist.ListRequest
	PlaylistRequest playlist.PlaylistRequest
	AddContent      playlist.PlaylistService_AddContentClient
	RemoveContent   playlist.PlaylistService_RemoveContentServer
)

type (
	Playlist playlist.Playlist
	List     playlist.ListResponse
)
