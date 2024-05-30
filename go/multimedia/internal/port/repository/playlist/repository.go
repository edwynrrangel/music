package playlist

import (
	"context"

	"github.com/edwynrrangel/grpc/go/multimedia/config"
	"github.com/edwynrrangel/grpc/go/multimedia/internal/domain/playlist"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	dbClient       *mongo.Client
	dbName         string
	collectionName string
}

func NewRepository(dbClient *mongo.Client, config *config.Config) playlist.Repository {
	return &repository{
		dbClient:       dbClient,
		dbName:         config.MongoDB.DbName,
		collectionName: config.MongoDB.PlaylistCollectionName,
	}
}

func (r *repository) Add(ctx context.Context, playlist playlist.PlayList) error {
	_, err := r.dbClient.Database(r.dbName).Collection(r.collectionName).InsertOne(ctx, playlist)
	return err
}

func (r *repository) Update(ctx context.Context, userID, playlistID string, playlist playlist.PlayList) error {
	objID, _ := primitive.ObjectIDFromHex(playlistID)
	filter := bson.M{"_id": objID, "user_id": userID}

	_, err := r.dbClient.Database(r.dbName).Collection(r.collectionName).ReplaceOne(ctx, filter, playlist)
	return err
}

func (r *repository) Remove(ctx context.Context, userID, playlistID string) error {
	objID, _ := primitive.ObjectIDFromHex(playlistID)
	filter := bson.M{"_id": objID, "user_id": userID}

	_, err := r.dbClient.Database(r.dbName).Collection(r.collectionName).DeleteOne(ctx, filter)
	return err
}

func (r *repository) Get(ctx context.Context, userID, playlistID string) (*playlist.PlayList, error) {
	objID, _ := primitive.ObjectIDFromHex(playlistID)
	filter := bson.M{"_id": objID, "user_id": userID}

	var playlist playlist.PlayList
	if err := r.dbClient.Database(r.dbName).Collection(r.collectionName).FindOne(ctx, filter).Decode(&playlist); err != nil {
		return nil, err
	}

	return &playlist, nil
}
