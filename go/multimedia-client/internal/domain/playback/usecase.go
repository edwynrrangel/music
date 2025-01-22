package playback

import (
	"context"
	"io"
	"log"

	"github.com/edwynrrangel/grpc/go/multimedia-client/internal/adapter/grpc/playback"
)

type usecase struct {
	playback playback.PlaybackServiceClient
}

// NewUseCase creates a new playback usecase
func NewUseCase(playback playback.PlaybackServiceClient) UseCase {
	return &usecase{
		playback: playback,
	}
}

func (u *usecase) Stream(ctx context.Context, data *ContentRequest, reactor Reactor) error {
	stream, err := u.playback.StreamSong(ctx, (*playback.ContentRequest)(data))
	if err != nil {
		return err
	}

	for {
		got, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error: %v", err)
			reactor.Publish(StreamEvent{Error: err})
			return err
		}
		reactor.Publish(StreamEvent{Data: got.Data})
	}

	reactor.Publish(StreamEvent{Done: true})
	return nil
}
