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
		Mode     string    `bson:"mode,omitempty"`
		Contents []Content `bson:"contents,omitempty"`
	}
	Playlists []Playlist
)
