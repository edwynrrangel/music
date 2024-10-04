package playlist

import "context"

type UseCase interface {
	Manage(ctx context.Context, request *PlaylistRequest) (*Playlist, error)
	Get(ctx context.Context, userID string, playlistID string) (*Playlist, error)
	List(ctx context.Context, userID string) ([]Playlist, error)
	Remove(ctx context.Context, userID string, playlistID string) error
	add(ctx context.Context, data *Playlist) (*Playlist, error)
	removeContent(ctx context.Context, data *Playlist) (*Playlist, error)
}

type Repository interface {
	Get(ctx context.Context, userID, playlistID string) (*Playlist, error)
	List(ctx context.Context, userID string) ([]Playlist, error)
	Add(ctx context.Context, playlist *Playlist) error
	Update(ctx context.Context, userID, playlistID string, data []Content) error
	Remove(ctx context.Context, userID, playlistID string) error
	RemoveContent(ctx context.Context, userID, playlistID string, contentID []string) error
}
