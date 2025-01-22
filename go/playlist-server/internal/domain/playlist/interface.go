package playlist

import (
	"context"

	"github.com/edwynrrangel/grpc/go/playlist-server/internal/shared"
)

type UseCase interface {
	Create(ctx context.Context, data *CreateRequest) (*Playlist, error)
	List(ctx context.Context, data ListRequest) (*shared.Paginated[Playlist], error)
	Get(ctx context.Context, userId string, playlistId string) (*Playlist, error)
	AddContent(ctx context.Context, userId string, playlistId string, content Content) (*Playlist, error)
	RemoveContent(ctx context.Context, userId string, playlistId string, contentId string) (*Playlist, error)
	PartyMode(ctx context.Context, data *PartyModeRequest) (*Playlist, error)
}

type Repository interface {
	Create(ctx context.Context, playlist *Playlist) error
	Count(ctx context.Context, userId, query string) (int64, error)
	List(ctx context.Context, userId, query string) (Playlists, error)
	Get(ctx context.Context, userId, playlistId string) (*Playlist, error)
	GetByType(ctx context.Context, plType string, playlistId string) (*Playlist, error)
	Update(ctx context.Context, userId, playlistId string, playlist *Playlist) error
	AddContent(ctx context.Context, playlistId string, content []Content) error
	RemoveContent(ctx context.Context, userId, playlistId, contentId string) error
}
