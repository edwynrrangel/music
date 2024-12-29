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
		playlist, err := a.usecase.AddContent(stream.Context(), data.UserId, data.PlaylistId, playlist.Content{
			ID:    data.ContentId,
			Order: data.Order,
		})
		if err != nil {
			return err
		}
		got = (*PlaylistEntity)(playlist).toResponse()
	}
}

func (a *adapter) RemoveContent(stream PlaylistService_RemoveContentServer) error {
	log.Printf("RemoveContent received request")
	var got *Playlist
	for {
		data, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(got)
		}
		if err != nil {
			return err
		}

		log.Printf("RemoveContent received request: %+v", data)
		playlist, err := a.usecase.RemoveContent(stream.Context(), data.UserId, data.PlaylistId, data.ContentId)
		if err != nil {
			return err
		}
		got = (*PlaylistEntity)(playlist).toResponse()
	}

}

/*
func (a *adapter) Manage(stream PlaylistService_ManageServer) error {
	log.Printf("Manage received request: %+v", stream)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		got, err := a.usecase.Manage(stream.Context(), req.toPlayListRequest())
		if err != nil {
			log.Printf("Error: %v", err)
			continue
		}

		stream.Send(convertToPlaylistResponse(got))
	}
}
*/
