package playlist

import (
	"context"

	"github.com/edwynrrangel/grpc/go/playlist_server/internal/shared"
)

type UseCase interface {
	Create(ctx context.Context, data *CreateRequest) (*Playlist, error)
	List(ctx context.Context, data ListRequest) (*shared.Paginated[Playlist], error)
	Get(ctx context.Context, userId string, playlistId string) (*Playlist, error)
	AddContent(ctx context.Context, userId string, playlistId, contentId string) (*Playlist, error)
}

type Repository interface {
	Create(ctx context.Context, playlist *Playlist) error
	Count(ctx context.Context, userId, query string) (int64, error)
	List(ctx context.Context, userId, query string) (Playlists, error)
	Get(ctx context.Context, userId, playlistId string) (*Playlist, error)
	Update(ctx context.Context, userId, playlistId, content string) error
}
