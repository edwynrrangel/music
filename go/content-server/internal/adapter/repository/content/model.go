package content

type (
	Location struct {
		Bucket string
		Key    string
	}
	Locations []Location
	Content   struct {
		ID        string   `bson:"_id"`
		Title     string   `bson:"title"`
		Artists   []string `bson:"artists"`
		Album     string   `bson:"album"`
		Genre     string   `bson:"genre"`
		Duration  int32    `bson:"duration"`
		CoverURL  string   `bson:"cover_url"`
		Locations Locations
	}
	Contents []Content
)
