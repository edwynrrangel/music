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
		playlistReq := req.toPlayListRequest()
		switch req.Operation {
		case PlaylistRequest_ADD:
			err = a.usecase.Add(stream.Context(), playlistReq)
		case PlaylistRequest_REMOVE:
			err = a.usecase.RemoveContent(stream.Context(), playlistReq)
		case PlaylistRequest_GET:
		default:
			log.Printf("Unknown operation: %v", req.Operation)
			return nil
		}
		if err != nil {
			log.Printf("Error: %v", err)
			continue
		}

		playlist, err := a.usecase.Get(stream.Context(), playlistReq)
		if err != nil {
			log.Printf("Error: %v", err)
			continue
		}
		stream.Send(convertToPlayListResponse(playlist))
	}
}

func (a *adapter) List(ctx context.Context, req *ListPlaylistRequest) (*ListPlaylistResponse, error) {
	log.Printf("List received request: %+v", req)
	got, err := a.usecase.List(req.UserID)
	if err != nil {
		return nil, err
	}
	return convertToListPlaylistResponse(got), nil
}

func (a *adapter) Remove(ctx context.Context, req *RemovePlaylistRequest) (*RemovePlaylistResponse, error) {
	log.Printf("Remove received request: %+v", req)
	err := a.usecase.Remove(ctx, req.toRemovePlaylistRequest())
	if err != nil {
		return nil, err
	}
	return &RemovePlaylistResponse{
		Message: "Playlist removed",
	}, nil
}
