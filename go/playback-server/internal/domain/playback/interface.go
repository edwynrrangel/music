package playback

import (
	"context"
	"io"
)

type UseCase interface {
	StreamSong(ctx context.Context, contentId string) (io.ReadCloser, error)
}
