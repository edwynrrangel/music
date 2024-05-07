package config

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type App struct {
	Port string `env:"PORT" envDefault:"50051"`
}

type Config struct {
	App App
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
