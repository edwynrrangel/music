package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/edwynrrangel/music/go/multimedia_server/config"
	"github.com/edwynrrangel/music/go/multimedia_server/pkg/mongodb"
)

// GetMongoClient function returns a mongo client
func GetMongoClient(cfg config.Config) *mongo.Client {
	builder := mongodb.NewBuilder(
		cfg.MongoDB.Connection.Host,
		cfg.MongoDB.Connection.Port,
		cfg.MongoDB.Connection.Username,
		cfg.MongoDB.Connection.Password,
	).WithRetryWrites("false")
	if cfg.MongoDB.Connection.TLS {
		builder = builder.WithTLS(cfg.MongoDB.Connection.CA)
	}
	client, err := builder.GetClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %s", err.Error())
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %s", err.Error())
	}

	return client
}
