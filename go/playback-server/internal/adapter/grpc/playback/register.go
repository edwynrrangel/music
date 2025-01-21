package playback

import (
	"github.com/edwynrrangel/go-bucket"
	"google.golang.org/grpc"

	"github.com/edwynrrangel/music/go/playback-server/config"
	"github.com/edwynrrangel/music/go/playback-server/internal/adapter/grpc/content"
	domain "github.com/edwynrrangel/music/go/playback-server/internal/domain/playback"
)

// Register function registers the playback service to the grpc server
func Register(s *grpc.Server, conn *grpc.ClientConn, config *config.Config, bucketClient bucket.Strategy) {
	RegisterPlaybackServiceServer(
		s,
		NewAdapter(
			domain.NewUseCase(
				bucketClient,
				content.NewContentServiceClient(conn),
			),
		),
	)
}
