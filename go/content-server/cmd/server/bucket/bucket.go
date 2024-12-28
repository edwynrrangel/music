package bucket

import (
	"log"

	"github.com/edwynrrangel/music/go/multimedia_server/config"
	"github.com/edwynrrangel/music/go/multimedia_server/pkg/bucket"
)

// GetClient function returns a bucket client
func GetClient(cfg config.Config) bucket.Strategy {
	switch cfg.Bucket.Type {
	case "minio":
		builder := bucket.NewBuilder(
			cfg.MinIO.Endpoint,
			cfg.MinIO.ReadonlyAccessKey,
			cfg.MinIO.ReadonlySecretKey,
		)
		if cfg.MinIO.UseSSL {
			builder = builder.WithSSL()
		}
		client, err := builder.GetClient()
		if err != nil {
			log.Fatalf("Failed to connect to Bucket: %s", err.Error())
		}
		return client
	default:
		log.Fatalf("Unknown bucket: %s", cfg.Bucket.Type)
	}
	return nil
}
