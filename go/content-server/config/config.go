package config

import (
	"log"
	"os"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

// app struct holds the configuration for the app
type app struct {
	Port string `env:"PORT" envDefault:"50051"`
}

// mongoDBConn struct holds the configuration for the mongoDB connection
type mongoDBConn struct {
	Host     string `env:"MONGO_HOST" envDefault:"localhost"`
	Port     string `env:"MONGO_PORT" envDefault:"27017"`
	Username string `env:"MONGO_USERNAME"`
	Password string `env:"MONGO_PASSWORD"`
	TLS      bool   `env:"MONGO_TLS" envDefault:"false"`
	CA       string `env:"MONGO_TLS_CA"`
}

// mongoDB struct holds the configuration for the mongoDB
type mongoDB struct {
	Connection             mongoDBConn
	DbName                 string `env:"MONGO_DB_NAME" envDefault:"test"`
	ContentCollectionnName string `env:"MONGO_CONTENT_COLLECTION_NAME" envDefault:"contents"`
	PlaylistCollectionName string `env:"MONGO_PLAYLIST_COLLECTION_NAME" envDefault:"playlists"`
}

// minIO struct holds the configuration for the minIO
type minIO struct {
	Endpoint          string `env:"MINIO_ENDPOINT" envDefault:"localhost:9000"`
	ReadonlyAccessKey string `env:"MINIO_READONLY_ACCESS_KEY"`
	ReadonlySecretKey string `env:"MINIO_READONLY_SECRET_KEY"`
	UseSSL            bool   `env:"MINIO_SSL" envDefault:"false"`
}

// bucket struct holds the configuration for the bucket
type bucket struct {
	Type string `env:"BUCKET_TYPE"`
}

// Config struct holds the configuration for the app, mongoDB, minIO and bucket
type Config struct {
	App     app
	MongoDB mongoDB
	MinIO   minIO
	Bucket  bucket
}

// init function sets the log flags and loads the .env file
func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if os.Getenv("ENV") != "prod" {
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
