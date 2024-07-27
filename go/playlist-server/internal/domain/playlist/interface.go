package playlist

import "context"

type UseCase interface {
	Add(ctx context.Context, request *PlaylistRequest) error
	RemoveContent(ctx context.Context, request *PlaylistRequest) error
	Get(ctx context.Context, request *PlaylistRequest) (*PlaylistResponse, error)
	List(userID string) ([]Playlist, error)
	Remove(ctx context.Context, request *RemovePlaylistRequest) error
}

type Repository interface {
	Get(ctx context.Context, playListID, user string) (*Playlist, error)
	List(ctx context.Context, userID string) ([]Playlist, error)
	Add(ctx context.Context, playList *Playlist) error
	Update(ctx context.Context, playListID, userID string, data Content) error
	RemoveContent(ctx context.Context, playListID, userID, contentID string) error
	Remove(ctx context.Context, playListID, userID string) error
}
