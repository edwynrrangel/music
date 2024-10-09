package playlist

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/edwynrrangel/grpc/go/multimedia_client/config"
	adapterPlaylist "github.com/edwynrrangel/grpc/go/multimedia_client/internal/adapter/grpc/playlist"
	domainPlaylist "github.com/edwynrrangel/grpc/go/multimedia_client/internal/domain/playlist"
)

func RegisterRoutes(api *gin.RouterGroup, cfg *config.Config, conn *grpc.ClientConn) {
	h := NewHandler(
		domainPlaylist.NewUseCase(
			adapterPlaylist.NewPlaylistServiceClient(conn),
		),
	)
	group := api.Group("/playlists")
	group.GET("user/:user_id", h.List)
	group.DELETE(":id/user/:user_id", h.Remove)
	group.GET(":id/manage", h.Manage)
}
