package playlist

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/edwynrrangel/grpc/go/playlist_server/internal/domain/playlist"
)

type repository struct {
	dbClient       *mongo.Client
	dbName         string
	collectionName string
}

func NewRepository(dbClient *mongo.Client, dbName string, collectionName string) playlist.Repository {
	return &repository{
		dbClient:       dbClient,
		dbName:         dbName,
		collectionName: collectionName,
	}
}

func (r *repository) Get(ctx context.Context, userID, playlistID string) (*playlist.Playlist, error) {
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

// Update update playlist
func (r *repository) Update(ctx context.Context, userID, playlistID string, content []playlist.Content) error {
	objID, _ := primitive.ObjectIDFromHex(playlistID)
	filter := bson.M{"_id": objID, "user_id": userID}
	update := bson.M{"$push": bson.M{"contents": bson.M{"$each": content}}}

	_, err := r.dbClient.Database(r.dbName).Collection(r.collectionName).UpdateOne(ctx, filter, update)
	return err
}

// RemoveContent remove contents from playlist
func (r *repository) RemoveContent(ctx context.Context, userID, playlistID string, contentIDs []string) error {
	objID, _ := primitive.ObjectIDFromHex(playlistID)
	filter := bson.M{"_id": objID, "user_id": userID}
	update := bson.M{"$pull": bson.M{"contents": bson.M{"id": bson.M{"$in": contentIDs}}}}

	_, err := r.dbClient.Database(r.dbName).Collection(r.collectionName).UpdateOne(ctx, filter, update)
	return err
}

// Remove remove playlist
func (r *repository) Remove(ctx context.Context, userID, playlistID string) error {
	objID, _ := primitive.ObjectIDFromHex(playlistID)
	filter := bson.M{"_id": objID, "user_id": userID}

	_, err := r.dbClient.Database(r.dbName).Collection(r.collectionName).DeleteOne(ctx, filter)
	return err
}
