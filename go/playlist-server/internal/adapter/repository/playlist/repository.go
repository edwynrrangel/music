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

func (r *repository) Create(ctx context.Context, data *playlist.Playlist) error {
	result, err := r.dbClient.Database(r.dbName).Collection(r.collectionName).InsertOne(ctx, (*PlaylistEntity)(data).toModel())
	if err != nil {
		return err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		data.ID = oid.Hex()
	}

	return nil
}

func (r *repository) Count(ctx context.Context, userId, query string) (int64, error) {
	filter := bson.M{"user_id": userId}
	if query != "" {
		filter["name"] = bson.M{"$regex": query, "$options": "i"}
	}

	return r.dbClient.Database(r.dbName).Collection(r.collectionName).CountDocuments(ctx, filter)
}

func (r *repository) List(ctx context.Context, userId, query string) (playlist.Playlists, error) {
	filter := bson.M{"user_id": userId}
	if query != "" {
		filter["name"] = bson.M{"$regex": query, "$options": "i"}
	}
	cursor, err := r.dbClient.Database(r.dbName).Collection(r.collectionName).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var playlists Playlists
	if err := cursor.All(ctx, &playlists); err != nil {
		return nil, err
	}

	return playlists.toEntity(), nil
}

func (r *repository) Get(ctx context.Context, userId, playlistId string) (*playlist.Playlist, error) {
	objID, _ := primitive.ObjectIDFromHex(playlistId)
	filter := bson.M{"_id": objID, "user_id": userId}

	var playlist = new(Playlist)
	if err := r.dbClient.Database(r.dbName).Collection(r.collectionName).FindOne(ctx, filter).Decode(playlist); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return playlist.toEntity(), nil
}

func (r *repository) GetByType(ctx context.Context, mode string, playlistId string) (*playlist.Playlist, error) {
	objID, _ := primitive.ObjectIDFromHex(playlistId)
	filter := bson.M{"_id": objID, "mode": mode}

	var playlist = new(Playlist)
	if err := r.dbClient.Database(r.dbName).Collection(r.collectionName).FindOne(ctx, filter).Decode(playlist); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return playlist.toEntity(), nil
}

func (r *repository) Update(ctx context.Context, userId, playlistId string, playlist *playlist.Playlist) error {
	objID, _ := primitive.ObjectIDFromHex(playlistId)
	filter := bson.M{"_id": objID, "user_id": userId}
	update := bson.M{
		"$set": bson.M{
			"name":     playlist.Name,
			"contents": playlist.Data,
		},
	}

	_, err := r.dbClient.Database(r.dbName).Collection(r.collectionName).UpdateOne(ctx, filter, update)
	return err
}

func (r *repository) AddContent(ctx context.Context, playlistId string, content []playlist.Content) error {
	objID, _ := primitive.ObjectIDFromHex(playlistId)
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"contents": content}}

	_, err := r.dbClient.Database(r.dbName).Collection(r.collectionName).UpdateOne(ctx, filter, update)
	return err
}

func (r *repository) RemoveContent(ctx context.Context, userId, playlistId, contentId string) error {
	objID, _ := primitive.ObjectIDFromHex(playlistId)
	filter := bson.M{"_id": objID, "user_id": userId}
	update := bson.M{"$pull": bson.M{"contents": bson.M{"id": contentId}}}

	_, err := r.dbClient.Database(r.dbName).Collection(r.collectionName).UpdateOne(ctx, filter, update)
	return err
}
