package multimedia

type (
	Content struct {
		ID      string `bson:"_id"`
		Title   string `bson:"title"`
		Genre   string `bson:"genre"`
		Creator string `bson:"creator"`
		Bucket  string `bson:"bucket"`
	}
)
