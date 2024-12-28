package content

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/edwynrrangel/music/go/multimedia_server/config"
	"github.com/edwynrrangel/music/go/multimedia_server/internal/domain/content"
)

type repository struct {
	dbClient       *mongo.Client
	dbName         string
	collectionName string
}

func NewRepository(dbClient *mongo.Client, config *config.Config) content.Repository {
	return &repository{
		dbClient:       dbClient,
		dbName:         config.MongoDB.DbName,
		collectionName: config.MongoDB.ContentCollectionnName,
	}
}

func (r *repository) Count(ctx context.Context, query string) (int64, error) {
	filter := bson.M{"$or": []bson.M{
		{"title": bson.M{"$regex": query, "$options": "i"}},
		{"genre": bson.M{"$regex": query, "$options": "i"}},
		{"artists": bson.M{"$elemMatch": bson.M{"$regex": query, "$options": "i"}}},
	}}

	count, err := r.dbClient.Database(r.dbName).Collection(r.collectionName).CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *repository) Search(ctx context.Context, query string) ([]content.Content, error) {
	filter := bson.M{"$or": []bson.M{
		{"title": bson.M{"$regex": query, "$options": "i"}},
		{"genre": bson.M{"$regex": query, "$options": "i"}},
		{"artists": bson.M{"$elemMatch": bson.M{"$regex": query, "$options": "i"}}},
	}}

	cursor, err := r.dbClient.Database(r.dbName).Collection(r.collectionName).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var contents Contents
	if err := cursor.All(ctx, &contents); err != nil {
		return nil, err
	}

	return contents.toArrayEntity(), nil
}
