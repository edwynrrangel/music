package content

import (
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"

	"github.com/edwynrrangel/grpc/go/multimedia/config"
	"github.com/edwynrrangel/grpc/go/multimedia/internal/content"
	"github.com/edwynrrangel/grpc/go/multimedia/internal/proto"
	"github.com/edwynrrangel/grpc/go/multimedia/pkg/bucket"
)

func Register(s *grpc.Server, config *config.Config, dbClient *mongo.Client, bucketClient bucket.Strategy) {
	proto.RegisterMultimediaServiceServer(
		s,
		NewAdapter(
			content.NewUseCase(
				content.NewRepository(dbClient, config),
				bucketClient,
			),
		),
	)
}
