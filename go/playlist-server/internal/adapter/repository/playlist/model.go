package playlist

type (
	Playlist struct {
		ID       string   `bson:"_id,omitempty"`
		UserId   string   `bson:"user_id"`
		Name     string   `bson:"name"`
		Contents []string `bson:"contents,omitempty"`
	}
	Playlists []Playlist
)
