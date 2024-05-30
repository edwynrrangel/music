package playlist

type (
	Content struct {
		ID      string `bson:"_id"`
		Title   string `bson:"title"`
		Genre   string `bson:"genre"`
		Creator string `bson:"creator"`
		Bucket  string `bson:"bucket"`
	}

	PlayList struct {
		ID      string    `bson:"_id"`
		UserID  string    `bson:"user_id"`
		Name    string    `bson:"name"`
		Content []Content `bson:"content"`
	}
)
