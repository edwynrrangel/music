package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/edwynrrangel/grpc/go/multimedia_client/config"
	"github.com/edwynrrangel/grpc/go/multimedia_client/internal/adapter/api/content"
	"github.com/edwynrrangel/grpc/go/multimedia_client/internal/adapter/api/playlist"
	"github.com/edwynrrangel/grpc/go/multimedia_client/internal/adapter/api/resource"
	"github.com/edwynrrangel/grpc/go/multimedia_client/internal/adapter/grpc/client"
)

// Run function starts the server
func Run() {
	cfg := config.New()
	contentServerConn, err := client.GetConn(
		cfg.Grpc.ContentServerUri,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	defer contentServerConn.Close()

	playlistServerConn, err := client.GetConn(
		cfg.Grpc.PlaylistServerUri,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	defer playlistServerConn.Close()

	r := gin.Default()
	resource.RegisterRoutes(r)

	api := r.Group("/api")
	content.RegisterRoutes(api, cfg, contentServerConn)
	playlist.RegisterRoutes(api, cfg, playlistServerConn)

	r.Run(fmt.Sprintf(":%s", cfg.App.Port))
}
