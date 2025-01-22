package content

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	grpcContent "github.com/edwynrrangel/grpc/go/multimedia-client/internal/adapter/grpc/content"
	domainContent "github.com/edwynrrangel/grpc/go/multimedia-client/internal/domain/content"
)

func RegisterRoutes(api *gin.RouterGroup, conn *grpc.ClientConn) {
	h := NewHandler(
		domainContent.NewUseCase(
			grpcContent.NewContentServiceClient(conn),
		),
	)

	group := api.Group("/content")
	group.GET("", h.Search)
	group.GET(":id", h.Get)
}
