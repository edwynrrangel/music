package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/edwynrrangel/grpc/go/playlist_server/cmd/server/database"
	"github.com/edwynrrangel/grpc/go/playlist_server/config"
	"github.com/edwynrrangel/grpc/go/playlist_server/internal/port/grpc/client"
	"github.com/edwynrrangel/grpc/go/playlist_server/internal/port/grpc/playlist"
)

// Run function starts the server
func Run() {
	cfg := config.New()

	dbClient := database.GetMongoClient(cfg)
	defer dbClient.Disconnect(context.Background())

	grpcConn, err := client.GetConn(
		cfg.Grpc.ContentServerUri,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	defer grpcConn.Close()

	grpcServer := grpc.NewServer()
	playlist.Register(grpcServer, &cfg, dbClient, grpcConn)

	log.Printf("Starting server on port %s", cfg.App.Port)
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.App.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}
