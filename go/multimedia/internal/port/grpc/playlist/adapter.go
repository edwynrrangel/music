package playlist

import (
	"io"
	"log"

	"github.com/edwynrrangel/grpc/go/multimedia/internal/domain/playlist"
)

type adapter struct {
	UnimplementedPlaylistServiceServer
	usecase playlist.UseCase
}

func NewAdapter(usecase playlist.UseCase) *adapter {
	return &adapter{usecase: usecase}
}

func (a *adapter) Manage(stream PlaylistService_ManageServer) error {
	log.Printf("Received request: %+v", stream)
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
			err = a.usecase.Remove(stream.Context(), playlistReq)
		case PlaylistRequest_GET:
		default:
			return nil
		}
		if err != nil {
			log.Printf("Error: %v", err)
		}

		playlist, err := a.usecase.Get(stream.Context(), playlistReq)
		if err != nil {
			return err
		}
		stream.Send(convertToPlayListResponse(playlist))
	}
}
