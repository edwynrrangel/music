package multimedia

import (
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"

	"github.com/edwynrrangel/grpc/go/multimedia/config"
	"github.com/edwynrrangel/grpc/go/multimedia/internal/multimedia"
	"github.com/edwynrrangel/grpc/go/multimedia/internal/proto"
	"github.com/edwynrrangel/grpc/go/multimedia/pkg/bucket"
)

func Register(s *grpc.Server, config *config.Config, dbClient *mongo.Client, bucketClient bucket.Strategy) {
	proto.RegisterMultimediaServiceServer(
		s,
		NewAdapter(
			multimedia.NewUseCase(
				multimedia.NewRepository(dbClient, config),
				bucketClient,
			),
		),
	)
}
