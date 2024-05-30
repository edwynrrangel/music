package playlist

import (
	"io"
	"log"

	"github.com/edwynrrangel/grpc/go/multimedia/internal/domain/playlist"
)

type adapter struct {
	UnimplementedPlayListServiceServer
	usecase playlist.UseCase
}

func NewAdapter(usecase playlist.UseCase) *adapter {
	return &adapter{usecase: usecase}
}

func (a *adapter) ManagePlayList(stream PlayListService_ManageServer) error {
	log.Printf("Received request: %+v", stream)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		switch req.Operation {
		case PlaylistRequest_ADD:
			a.usecase.Add(stream.Context(), *req.toPlayListRequest())
		case PlaylistRequest_REMOVE:
			a.usecase.Remove(stream.Context(), *req.toPlayListRequest())
		case PlaylistRequest_GET:
			playlist, err := a.usecase.Get(stream.Context(), *req.toPlayListRequest())
			if err != nil {
				return err
			}
			return stream.Send(convertToPlayListResponse(playlist))
		default:
			return nil
		}
	}
}
