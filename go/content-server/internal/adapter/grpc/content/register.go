package content

import (
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"

	"github.com/edwynrrangel/music/go/content-server/config"
	contentRepo "github.com/edwynrrangel/music/go/content-server/internal/adapter/repository/content"
	contentDomain "github.com/edwynrrangel/music/go/content-server/internal/domain/content"
)

// Register function registers the content service to server gRPC
func Register(s *grpc.Server, config *config.Config, dbClient *mongo.Client) {
	RegisterContentServiceServer(
		s,
		NewAdapter(
			contentDomain.NewUseCase(
				contentRepo.NewRepository(dbClient, config),
			),
		),
	)
}
