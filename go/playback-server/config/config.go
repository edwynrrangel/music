package config

import (
	"log"
	"os"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

// app struct holds the configuration for the app
type app struct {
	Port string `env:"PORT" envDefault:"50053"`
}

// minIO struct holds the configuration for the minIO
type minIO struct {
	Endpoint          string `env:"MINIO_ENDPOINT" envDefault:"localhost:9000"`
	ReadonlyAccessKey string `env:"MINIO_READONLY_ACCESS_KEY"`
	ReadonlySecretKey string `env:"MINIO_READONLY_SECRET_KEY"`
	UseSSL            bool   `env:"MINIO_SSL" envDefault:"false"`
}

type bucket struct {
	Type string `env:"BUCKET_TYPE"`
}

// grpc struct holds the configuration for the grpc server
type grpc struct {
	ContentServerUri  string `env:"CONTENT_SERVER_URI" envDefault:"localhost:50051"`
	PlaylistServerUri string `env:"PLAYLIST_SERVER_URI" envDefault:"localhost:50052"`
}

// Config struct holds the configuration for the app, mongoDB, minIO and bucket
type Config struct {
	App    app
	MinIO  minIO
	Bucket bucket
	Grpc   grpc
}

// init function sets the log flags and loads the .env file
func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if os.Getenv("ENV") != "" {
		return
	}
	if err := godotenv.Load(); err != nil {
		log.Println("Warning loading .env file")
	}
}

// New function creates a new Config instance and parses the environment variables
func New() Config {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalf("Failed to parse config: %s", err.Error())
	}
	return cfg
}
