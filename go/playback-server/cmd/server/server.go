package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/edwynrrangel/music/go/playback-server/cmd/server/bucket"
	"github.com/edwynrrangel/music/go/playback-server/config"
	"github.com/edwynrrangel/music/go/playback-server/internal/adapter/grpc/playback"
)

// Run function starts the server
func Run() {
	cfg := config.New()

	bucketClient := bucket.GetClient(cfg)
	contentConn, err := grpc.NewClient(
		cfg.Grpc.ContentServerUri,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect to content server: %s", err.Error())
	}

	grpcServer := grpc.NewServer()
	playback.Register(grpcServer, contentConn, &cfg, bucketClient)

	log.Printf("Starting server on port %s", cfg.App.Port)
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.App.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}
