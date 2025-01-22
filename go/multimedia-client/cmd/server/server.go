package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/edwynrrangel/grpc/go/multimedia-client/config"
	"github.com/edwynrrangel/grpc/go/multimedia-client/internal/adapter/api/content"
	"github.com/edwynrrangel/grpc/go/multimedia-client/internal/adapter/api/playback"
	"github.com/edwynrrangel/grpc/go/multimedia-client/internal/adapter/api/playlist"
	"github.com/edwynrrangel/grpc/go/multimedia-client/internal/adapter/api/resource"
	"github.com/edwynrrangel/grpc/go/multimedia-client/internal/adapter/grpc/client"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Run function starts the server
func Run() {
	cfg := config.New()
	contentServerConn, err := client.GetConn(
		cfg.Grpc.ContentServerUri,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer contentServerConn.Close()

	playlistServerConn, err := client.GetConn(
		cfg.Grpc.PlaylistServerUri,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer playlistServerConn.Close()

	playbackServerConn, err := client.GetConn(
		cfg.Grpc.PlaybackServerUri,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	resource.RegisterRoutes(r)

	api := r.Group("/api")
	content.RegisterRoutes(api, contentServerConn)
	playlist.RegisterRoutes(api, playlistServerConn)
	playback.RegisterRoutes(api, playbackServerConn)

	r.Run(fmt.Sprintf(":%s", cfg.App.Port))
}
