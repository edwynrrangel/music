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
	DbName                 string `env:"MONGO_DB_NAME" envDefault:"multimedia"`
	PlaylistCollectionName string `env:"MONGO_PLAYLIST_COLLECTION_NAME" envDefault:"playlists"`
}

// Config struct holds the configuration for the app, mongoDB, minIO and bucket
type Config struct {
	App     app
	MongoDB mongoDB
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
