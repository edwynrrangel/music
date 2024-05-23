package multimedia

import (
	"context"

	"github.com/edwynrrangel/grpc/go/multimedia/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	dbClient       *mongo.Client
	dbName         string
	collectionName string
}

func NewRepository(dbClient *mongo.Client, config *config.Config) Repository {
	return &repository{
		dbClient:       dbClient,
		dbName:         config.MongoDB.DbName,
		collectionName: config.MongoDB.CollectionName,
	}
}

func (r *repository) SearchContent(ctx context.Context, query string) ([]Content, error) {
	filter := bson.M{"$or": []bson.M{
		{"title": bson.M{"$regex": query, "$options": "i"}},
		{"genre": bson.M{"$regex": query, "$options": "i"}},
		{"creator": bson.M{"$regex": query, "$options": "i"}},
	}}

	cursor, err := r.dbClient.Database(r.dbName).Collection(r.collectionName).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var contents []Content
	if err := cursor.All(ctx, &contents); err != nil {
		return nil, err
	}

	return contents, nil
}
