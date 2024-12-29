package playlist

type (
	Content struct {
		ID    string `bson:"id"`
		Order int32  `bson:"order"`
	}
	Playlist struct {
		ID       string    `bson:"_id,omitempty"`
		UserId   string    `bson:"user_id"`
		Name     string    `bson:"name"`
		Contents []Content `bson:"contents,omitempty"`
	}
	Playlists []Playlist
)
