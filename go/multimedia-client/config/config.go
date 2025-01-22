package config

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

// app struct holds the configuration for the app
type app struct {
	Port string `env:"PORT" envDefault:"3000"`
}

// grpc struct holds the configuration for the gRPC server
type grpc struct {
	ContentServerUri  string `env:"CONTENT_SERVER_URI" envDefault:"localhost:50051"`
	PlaylistServerUri string `env:"PLAYLIST_SERVER_URI" envDefault:"localhost:50052"`
	PlaybackServerUri string `env:"PLAYBACK_SERVER_URI" envDefault:"localhost:50053"`
}

// Config struct holds the configuration
type Config struct {
	App  app
	Grpc grpc
}

// init function sets the log flags and loads the .env file
func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err := godotenv.Load(); err != nil {
		log.Println("Warning loading .env file")
	}
}

// New function creates a new Config instance and parses the environment variables
func New() *Config {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalf("Failed to parse config: %s", err.Error())
	}
	return &cfg
}
