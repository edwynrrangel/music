package main

import (
	"fmt"
	"log"
	"net"

	"github.com/edwynrrangel/grpc/go/multimedia/cmd/grpc/multimedia"
	"github.com/edwynrrangel/grpc/go/multimedia/config"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.New()
	log.Printf("Starting server on port %s", cfg.App.Port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.App.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	multimedia.Register(grpcServer, &cfg)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}
