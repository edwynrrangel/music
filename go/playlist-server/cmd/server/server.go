package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/edwynrrangel/grpc/go/playlist-server/cmd/server/database"
	"github.com/edwynrrangel/grpc/go/playlist-server/config"
	"github.com/edwynrrangel/grpc/go/playlist-server/internal/adapter/grpc/playlist"
)

// Run function starts the server
func Run() {
	cfg := config.New()

	dbClient := database.GetMongoClient(cfg)
	defer dbClient.Disconnect(context.Background())

	grpcServer := grpc.NewServer()
	playlist.Register(grpcServer, &cfg, dbClient)

	log.Printf("Starting server on port %s", cfg.App.Port)
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.App.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}
