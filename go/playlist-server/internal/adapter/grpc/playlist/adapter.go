package playlist

import (
	"context"
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



func (a *adapter) Remove(ctx context.Context, req *RemovePlaylistRequest) (*RemovePlaylistResponse, error) {
	log.Printf("Remove received request: %+v", req)
	err := a.usecase.Remove(ctx, req.UserID, req.Id)
	if err != nil {
		return nil, err
	}
	return &RemovePlaylistResponse{
		Message: "Playlist removed",
	}, nil
} */
