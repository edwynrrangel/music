package content

import (
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"

	"github.com/edwynrrangel/music/go/multimedia_server/config"
	contentRepo "github.com/edwynrrangel/music/go/multimedia_server/internal/adapter/repository/content"
	contentDomain "github.com/edwynrrangel/music/go/multimedia_server/internal/domain/content"
	"github.com/edwynrrangel/music/go/multimedia_server/pkg/bucket"
)

// Register function registers the content service to server gRPC
func Register(s *grpc.Server, config *config.Config, dbClient *mongo.Client, bucketClient bucket.Strategy) {
	RegisterContentServiceServer(
		s,
		NewAdapter(
			contentDomain.NewUseCase(
				bucketClient,
				contentRepo.NewRepository(dbClient, config),
			),
		),
	)
}
