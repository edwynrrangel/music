package playlist

import (
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"

	"github.com/edwynrrangel/grpc/go/playlist_server/config"
	playListDomain "github.com/edwynrrangel/grpc/go/playlist_server/internal/domain/playlist"
	"github.com/edwynrrangel/grpc/go/playlist_server/internal/port/grpc/content"
	playListRepo "github.com/edwynrrangel/grpc/go/playlist_server/internal/port/repository/playlist"
)

func Register(s *grpc.Server, config *config.Config, dbClient *mongo.Client, conn *grpc.ClientConn) {
	RegisterPlaylistServiceServer(
		s,
		NewAdapter(
			playListDomain.NewUseCase(
				playListRepo.NewRepository(dbClient, config),
				content.NewAdapter(content.NewContentServiceClient(conn)),
			),
		),
	)
}
