package playlist

import (
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"

	"github.com/edwynrrangel/grpc/go/playlist_server/config"
	playListRepo "github.com/edwynrrangel/grpc/go/playlist_server/internal/adapter/repository/playlist"
	playListDomain "github.com/edwynrrangel/grpc/go/playlist_server/internal/domain/playlist"
)

func Register(s *grpc.Server, config *config.Config, dbClient *mongo.Client) {
	RegisterPlaylistServiceServer(
		s,
		NewAdapter(
			playListDomain.NewUseCase(
				playListRepo.NewRepository(dbClient, config.MongoDB.DbName, config.MongoDB.PlaylistCollectionName),
			),
		),
	)
}
