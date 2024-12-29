package playlist

import (
	"context"
	"io"
	"log"

	"github.com/edwynrrangel/grpc/go/playlist_server/internal/domain/playlist"
)

type adapter struct {
	UnimplementedPlaylistServiceServer
	usecase playlist.UseCase
}

func NewAdapter(usecase playlist.UseCase) *adapter {
	return &adapter{usecase: usecase}
}

func (a *adapter) Create(ctx context.Context, data *CreateRequest) (*Playlist, error) {
	log.Printf("Create received request: %+v", data)
	got, err := a.usecase.Create(ctx, data.toRequest())
	if err != nil {
		return nil, err
	}

	return (*PlaylistEntity)(got).toResponse(), nil
}

func (a *adapter) List(ctx context.Context, data *ListRequest) (*ListResponse, error) {
	log.Printf("List received request: %+v", data)
	paginated, err := a.usecase.List(ctx, *data.toRequest())
	if err != nil {
		return nil, err
	}

	return (*Paginated)(paginated).toResponse(), nil
}

func (a *adapter) Get(ctx context.Context, data *PlaylistRequest) (*Playlist, error) {
	log.Printf("Get received request: %+v", data)
	got, err := a.usecase.Get(ctx, data.UserId, data.PlaylistId)
	if err != nil {
		return nil, err
	}

	return (*PlaylistEntity)(got).toResponse(), nil
}

func (a *adapter) AddContent(stream PlaylistService_AddContentServer) error {
	log.Printf("AddContent received request")
	var got *Playlist
	for {
		data, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(got)
		}
		if err != nil {
			return err
		}

		log.Printf("AddContent received request: %+v", data)
		playlist, err := a.usecase.AddContent(stream.Context(), data.UserId, data.PlaylistId, data.ContentId)
		if err != nil {
			return err
		}
		got = (*PlaylistEntity)(playlist).toResponse()
	}
}
