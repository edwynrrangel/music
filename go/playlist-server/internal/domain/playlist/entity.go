package playlist

type (
	Content struct {
		ID       string `bson:"id"`
		Title    string `bson:"title"`
		Creator  string `bson:"creator"`
		Genre    string `bson:"genre"`
		Duration string `bson:"duration"`
	}

	Playlist struct {
		ID      string    `bson:"_id,omitempty"`
		UserID  string    `bson:"user_id"`
		Name    string    `bson:"name"`
		Content []Content `bson:"content"`
	}
)
