package playback

import "github.com/edwynrrangel/grpc/go/multimedia-client/internal/adapter/grpc/playback"

type (
	ContentRequest playback.ContentRequest
	StreamEvent    struct {
		Data  []byte
		Error error
		Done  bool
	}
)
