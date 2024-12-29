package playlist

import (
	"context"

	"github.com/edwynrrangel/grpc/go/playlist_server/internal/shared"
)

type UseCase interface {
	Create(ctx context.Context, data *CreateRequest) (*Playlist, error)
	List(ctx context.Context, data ListRequest) (*shared.Paginated[Playlist], error)
	Get(ctx context.Context, userID string, playlistID string) (*Playlist, error)
	/*
		Manage(ctx context.Context, request *PlaylistRequest) (*Playlist, error)
		Remove(ctx context.Context, userID string, playlistID string) error
		add(ctx context.Context, data *Playlist) (*Playlist, error)
		removeContent(ctx context.Context, data *Playlist) (*Playlist, error)
	*/
}

type Repository interface {
	Create(ctx context.Context, playlist *Playlist) error
	Count(ctx context.Context, userID, query string) (int64, error)
	List(ctx context.Context, userID, query string) (Playlists, error)
	Get(ctx context.Context, userID, playlistID string) (*Playlist, error)
	/*

		Update(ctx context.Context, userID, playlistID string, data []Content) error
		Remove(ctx context.Context, userID, playlistID string) error
		RemoveContent(ctx context.Context, userID, playlistID string, contentID []string) error
	*/
}
