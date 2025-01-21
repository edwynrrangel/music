package playback

import (
	"context"
	"io"

	"github.com/edwynrrangel/go-bucket"
	"github.com/edwynrrangel/music/go/playback-server/internal/adapter/grpc/content"
)

type usecase struct {
	minIO   bucket.Strategy
	content content.ContentServiceClient
}

func NewUseCase(bucket bucket.Strategy, content content.ContentServiceClient) UseCase {
	return &usecase{
		minIO:   bucket,
		content: content,
	}
}

func (u *usecase) StreamSong(ctx context.Context, contentId string) (io.ReadCloser, error) {
	content, err := u.content.Get(ctx, &content.ContentRequest{Id: contentId})
	if err != nil {
		return nil, err
	}

	return u.minIO.DownloadFile(ctx, content.Locations[0].Bucket, content.Locations[0].Key)
}
