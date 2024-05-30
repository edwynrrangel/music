package playlist

import "context"

type UseCase interface {
	Add(ctx context.Context, request *PlayListRequest) error
	Remove(ctx context.Context, request *PlayListRequest) error
	Get(ctx context.Context, request *PlayListRequest) (*Playlist, error)
}

type Repository interface {
	Get(ctx context.Context, playListID, user string) (*Playlist, error)
	GetAll(ctx context.Context, userID string) ([]Playlist, uint32, error)
	Add(ctx context.Context, playList *Playlist) error
	Update(ctx context.Context, playListID, userID string, data Content) error
	RemoveContent(ctx context.Context, playListID, userID, contentID string) error
}
