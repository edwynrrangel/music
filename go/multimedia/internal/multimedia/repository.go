package multimedia

import (
	"github.com/edwynrrangel/grpc/go/multimedia/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct{}

func NewRepository(dbClient *mongo.Client, config *config.Config) Repository {
	return &repository{}
}

func (r *repository) SearchContent(query string) ([]Content, error) {
	contents := []Content{
		{ID: "1", Title: "Example Title 1", Genre: "Action", Creator: "Director A", Url: "http://example.com/1"},
		{ID: "2", Title: "Example Title 2", Genre: "Comedy", Creator: "Director B", Url: "http://example.com/2"},
	}

	return contents, nil
}
