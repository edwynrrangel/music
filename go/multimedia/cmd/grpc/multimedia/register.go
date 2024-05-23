package multimedia

import (
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"

	"github.com/edwynrrangel/grpc/go/multimedia/config"
	"github.com/edwynrrangel/grpc/go/multimedia/internal/multimedia"
	"github.com/edwynrrangel/grpc/go/multimedia/internal/proto"
)

func Register(s *grpc.Server, config *config.Config, dbClient *mongo.Client) {
	proto.RegisterMultimediaServiceServer(
		s,
		NewAdapter(
			multimedia.NewUseCase(
				multimedia.NewRepository(dbClient, config),
			),
		),
	)
}
