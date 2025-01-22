package playlist

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	grpcPlaylist "github.com/edwynrrangel/grpc/go/multimedia-client/internal/adapter/grpc/playlist"
	domainPlaylist "github.com/edwynrrangel/grpc/go/multimedia-client/internal/domain/playlist"
)

func RegisterRoutes(api *gin.RouterGroup, conn *grpc.ClientConn) {
	h := NewHandler(
		domainPlaylist.NewUseCase(
			grpcPlaylist.NewPlaylistServiceClient(conn),
		),
	)

	group := api.Group("/playlist")
	group.POST("", h.Create)
	group.GET("user/:user_id/list", h.List)
	group.GET(":id/user/:user_id", h.Get)
	// group.DELETE(":id/user/:user_id", h.Remove)
	// group.GET(":id/manage", h.Manage)
}
