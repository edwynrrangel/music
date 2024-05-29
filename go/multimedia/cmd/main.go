package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/edwynrrangel/grpc/go/multimedia/cmd/database"
	"github.com/edwynrrangel/grpc/go/multimedia/config"
)

func main() {
	cfg := config.New()

	dbClient := database.GetMongoClient(cfg)
	defer dbClient.Disconnect(context.Background())

	// bucketClient := bucket.GetClient(cfg)

	log.Printf("Starting server on port %s", cfg.App.Port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.App.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	// multimedia.Register(grpcServer, &cfg, dbClient, bucketClient)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}
