package config

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type app struct {
	Port string `env:"PORT" envDefault:"50051"`
}

type mongoDBConn struct {
	Host     string `env:"MONGO_HOST" envDefault:"localhost"`
	Port     string `env:"MONGO_PORT" envDefault:"27017"`
	Username string `env:"MONGO_USERNAME"`
	Password string `env:"MONGO_PASSWORD"`
	TLS      bool   `env:"MONGO_TLS" envDefault:"false"`
	CA       string `env:"MONGO_TLS_CA"`
}
type mongoDB struct {
	Connection     mongoDBConn
	DbName         string `env:"MONGO_DB_NAME" envDefault:"test"`
	CollectionName string `env:"MONGO_COLLECTION_NAME" envDefault:"multimedia"`
}
type minIO struct {
	Endpoint          string `env:"MINIO_ENDPOINT" envDefault:"localhost:9000"`
	ReadonlyAccessKey string `env:"MINIO_READONLY_ACCESS_KEY"`
	ReadonlySecretKey string `env:"MINIO_READONLY_SECRET_KEY"`
	UseSSL            bool   `env:"MINIO_SSL" envDefault:"false"`
}
type bucket struct {
	Type string `env:"BUCKET_TYPE"`
	Name string `env:"BUCKET_NAME"`
}
type Config struct {
	App     app
	MongoDB mongoDB
	MinIO   minIO
	Bucket  bucket
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err := godotenv.Load(); err != nil {
		log.Println("Warning loading .env file")
	}
}

func New() Config {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalf("Failed to parse config: %s", err.Error())
	}
	return cfg
}
