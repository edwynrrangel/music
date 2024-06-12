package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/edwynrrangel/grpc/go/multimedia_client/config"
	"github.com/edwynrrangel/grpc/go/multimedia_client/internal/port/api/content"
	"github.com/edwynrrangel/grpc/go/multimedia_client/internal/port/api/resource"
	"github.com/edwynrrangel/grpc/go/multimedia_client/internal/port/grpc/client"
)

// Run function starts the server
func Run() {
	cfg := config.New()
	grpcConn, err := client.GetConn(
		cfg.Grpc.ContentServerUri,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	defer grpcConn.Close()

	r := gin.Default()
	resource.RegisterRoutes(r)

	api := r.Group("/api")
	content.RegisterRoutes(api, cfg, grpcConn)

	r.Run(fmt.Sprintf(":%s", cfg.App.Port))
}
