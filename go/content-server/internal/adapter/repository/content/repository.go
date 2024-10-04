package content

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/edwynrrangel/grpc/go/multimedia_server/config"
	"github.com/edwynrrangel/grpc/go/multimedia_server/internal/domain/content"
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

func (r *repository) Get(ctx context.Context, id string) (*content.Content, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}

	var content *Content
	if err := r.dbClient.Database(r.dbName).Collection(r.collectionName).FindOne(ctx, filter).Decode(&content); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return content.toEntity(), nil
}

func (r *repository) Search(ctx context.Context, query string) ([]content.Content, error) {
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

	return toArrayEntity(contents), nil
}
