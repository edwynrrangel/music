package playlist

import (
	"context"

	"github.com/edwynrrangel/grpc/go/multimedia-client/internal/adapter/grpc/playlist"
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

// Create creates a new playlist
func (u *usecase) Create(ctx context.Context, data *CreateRequest) (*Playlist, error) {
	got, err := u.playlist.Create(ctx, (*playlist.CreateRequest)(data))
	if err != nil {
		return nil, err
	}
	return (*Playlist)(got), nil
}

// List returns a list of playlists
func (u *usecase) List(ctx context.Context, data *ListRequest) (*List, error) {
	got, err := u.playlist.List(ctx, (*playlist.ListRequest)(data))
	if err != nil {
		return nil, err
	}
	return (*List)(got), nil
}

// Get returns a playlist
func (u *usecase) Get(ctx context.Context, data *PlaylistRequest) (*Playlist, error) {
	got, err := u.playlist.Get(ctx, (*playlist.PlaylistRequest)(data))
	if err != nil {
		return nil, err
	}
	return (*Playlist)(got), nil
}

/* func (u *usecase) Manage(ctx context.Context, comm Communication, id string) error {
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
			comm.WriteJSON(map[string]string{"error": "userID is required"})
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
} */
