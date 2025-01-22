package playback

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	grpcPlayback "github.com/edwynrrangel/grpc/go/multimedia-client/internal/adapter/grpc/playback"
	domainPlayback "github.com/edwynrrangel/grpc/go/multimedia-client/internal/domain/playback"
)

func RegisterRoutes(api *gin.RouterGroup, conn *grpc.ClientConn) {
	h := NewHandler(
		domainPlayback.NewUseCase(
			grpcPlayback.NewPlaybackServiceClient(conn),
		),
	)

	group := api.Group("/playback")
	group.GET(":content_id/stream", reactorMiddleware(), h.Stream)
}

func reactorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("reactor", domainPlayback.NewReactor())
		c.Next()
	}
}
