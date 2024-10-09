package playlist

import (
	"context"
	"io"

	"github.com/edwynrrangel/grpc/go/multimedia_client/internal/adapter/grpc/playlist"
)

type usecase struct {
	playlist playlist.PlaylistServiceClient
}

// NewUseCase creates a new playlist usecase
func NewUseCase(playlist playlist.PlaylistServiceClient) UseCase {
	return &usecase{
		playlist: playlist,
	}
}

// List returns a list of playlists
func (u *usecase) List(ctx context.Context, filters ListFilter) (*ListResponse, error) {
	got, err := u.playlist.List(ctx, &playlist.ListPlaylistRequest{
		UserID: filters.UserID,
	})
	if err != nil {
		return nil, err
	}
	return convertToListResponse(got), nil
}

func (u *usecase) Remove(ctx context.Context, params RemoveRequest) error {
	_, err := u.playlist.Remove(ctx, &playlist.RemovePlaylistRequest{
		Id:     params.ID,
		UserID: params.UserID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *usecase) Manage(ctx context.Context, comm Communication, id string) error {
	stream, err := u.playlist.Manage(ctx)
	if err != nil {
		return err
	}

	// Goroutine to handle receiving multiple responses from the server
	go func() {
		for {
			got, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				comm.WriteJSON(map[string]string{"error": err.Error()})
				return
			}

			comm.WriteJSON(got)
		}
	}()

	// handle sending multiple requests as they come from the client
	for {
		var bodyRequest playlist.PlaylistRequest
		if err := comm.ReadJSON(&bodyRequest); err != nil {
			comm.WriteJSON(map[string]string{"error": err.Error()})
			return err
		}

		if bodyRequest.UserID == "" {
			comm.WriteJSON(map[string]string{"error": "user_id is required"})
			break
		}

		if bodyRequest.Operation == playlist.Operation_CLOSED {
			break
		}

		bodyRequest.Id = id
		if err := stream.Send(&bodyRequest); err != nil {
			comm.WriteJSON(map[string]string{"error": err.Error()})
			return err
		}
	}

	return nil
}
