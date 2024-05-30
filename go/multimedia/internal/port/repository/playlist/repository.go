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

func (r *repository) Get(ctx context.Context, playlistID, userID string) (*playlist.Playlist, error) {
	objID, _ := primitive.ObjectIDFromHex(playlistID)
	filter := bson.M{"_id": objID, "user_id": userID}

	var playlist = new(playlist.Playlist)
	if err := r.dbClient.Database(r.dbName).Collection(r.collectionName).FindOne(ctx, filter).Decode(playlist); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return playlist, nil
}

func (r *repository) List(ctx context.Context, userID string) ([]playlist.Playlist, error) {
	filter := bson.M{"user_id": userID}
	cursor, err := r.dbClient.Database(r.dbName).Collection(r.collectionName).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var playlists []playlist.Playlist
	if err := cursor.All(ctx, &playlists); err != nil {
		return nil, err
	}

	return playlists, nil
}

func (r *repository) Add(ctx context.Context, playlist *playlist.Playlist) error {
	result, err := r.dbClient.Database(r.dbName).Collection(r.collectionName).InsertOne(ctx, playlist)
	if err != nil {
		return err
	}
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		playlist.ID = oid.Hex()
	}
	return nil
}

func (r *repository) Update(ctx context.Context, playlistID, userID string, content playlist.Content) error {
	objID, _ := primitive.ObjectIDFromHex(playlistID)
	filter := bson.M{"_id": objID, "user_id": userID}
	update := bson.M{"$push": bson.M{"content": content}}

	_, err := r.dbClient.Database(r.dbName).Collection(r.collectionName).UpdateOne(ctx, filter, update)
	return err
}

// Remove removes content from playlist
func (r *repository) RemoveContent(ctx context.Context, playlistID, userID, contentID string) error {
	objID, _ := primitive.ObjectIDFromHex(playlistID)
	filter := bson.M{"_id": objID, "user_id": userID}
	update := bson.M{"$pull": bson.M{"content": bson.M{"_id": contentID}}}

	_, err := r.dbClient.Database(r.dbName).Collection(r.collectionName).UpdateOne(ctx, filter, update)
	return err
}
