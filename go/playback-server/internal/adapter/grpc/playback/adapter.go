package playback

import (
	"io"
	"log"

	"github.com/edwynrrangel/music/go/playback-server/internal/domain/playback"
)

type adapter struct {
	UnimplementedPlaybackServiceServer
	useCase playback.UseCase
}

func NewAdapter(useCase playback.UseCase) *adapter {
	return &adapter{useCase: useCase}
}

func (a *adapter) StreamSong(req *ContentRequest, stream PlaybackService_StreamSongServer) error {
	log.Printf("Received request: %+v", req)
	file, err := a.useCase.StreamSong(stream.Context(), req.ContentId)
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := make([]byte, 64*1024)
	cont := 0
	for {

		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		if err := stream.Send(&Song{Data: buffer[:n]}); err != nil {
			return err
		}
		cont++
		log.Printf("chunk %d sent", cont)
	}

	return nil
}
