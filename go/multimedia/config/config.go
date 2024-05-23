package config

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type app struct {
	Port string `env:"PORT" envDefault:"50051"`
}

type mongoDB struct {
	Host     string `env:"MONGO_HOST" envDefault:"localhost"`
	Port     string `env:"MONGO_PORT" envDefault:"27017"`
	DbName   string `env:"MONGO_DB_NAME" envDefault:"test"`
	Username string `env:"MONGO_USERNAME"`
	Password string `env:"MONGO_PASSWORD"`
	TLS      bool   `env:"MONGO_TLS" envDefault:"false"`
	CA       string `env:"MONGO_TLS_CA"`
}

type Config struct {
	App     app
	MongoDB mongoDB
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
