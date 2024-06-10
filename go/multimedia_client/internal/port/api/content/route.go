package content

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/edwynrrangel/grpc/go/multimedia_client/config"
	"github.com/edwynrrangel/grpc/go/multimedia_client/internal/port/grpc/content"
)

func RegisterRoutes(api *gin.RouterGroup, cfg *config.Config, conn *grpc.ClientConn) {
	h := NewHandler(content.NewContentServiceClient(conn))
	group := api.Group("/contents")
	group.GET(":id", h.Get)
	group.GET("", h.Search)
	group.GET(":id/stream", h.Stream)
}
