package playlist

import (
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"

	"github.com/edwynrrangel/grpc/go/multimedia/config"
	playListDomain "github.com/edwynrrangel/grpc/go/multimedia/internal/domain/playlist"
	playListRepo "github.com/edwynrrangel/grpc/go/multimedia/internal/port/repository/playlist"
)

func Register(s *grpc.Server, config *config.Config, dbClient *mongo.Client) {
	RegisterPlayListServiceServer(
		s,
		NewAdapter(
			playListDomain.NewUseCase(
				playListRepo.NewRepository(dbClient, config),
			),
		),
	)
}
