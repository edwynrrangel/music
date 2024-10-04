package content

type (
	Content struct {
		ID       string `bson:"_id"`
		Title    string `bson:"title"`
		Creator  string `bson:"creator"`
		Genre    string `bson:"genre"`
		Duration string `bson:"duration"`
		Bucket   string `bson:"bucket"`
	}
)
