package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/edwynrrangel/grpc/go/multimedia/config"
	"github.com/edwynrrangel/grpc/go/multimedia/pkg/mongodb"
)

func GetMongoClient(cfg config.Config) *mongo.Client {
	mongoBuilder := mongodb.NewBuilder(cfg).WithRetryWrites("false")
	if cfg.MongoDB.TLS {
		mongoBuilder = mongoBuilder.WithTLS(cfg.MongoDB.CA)
	}
	client, err := mongoBuilder.GetClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %s", err.Error())
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %s", err.Error())
	}

	return client
}
