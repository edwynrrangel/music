package content

type (
	Location struct {
		Bucket string `bson:"bucket"`
		Key    string `bson:"key"`
	}
	Locations []Location
	Content   struct {
		ID        string    `bson:"_id,omitempty"`
		Title     string    `bson:"title"`
		Artists   []string  `bson:"artists"`
		Album     string    `bson:"album"`
		Genre     string    `bson:"genre"`
		Duration  int32     `bson:"duration"`
		CoverURL  string    `bson:"cover_url"`
		Locations Locations `bson:"locations"`
	}

	Contents []Content
)
