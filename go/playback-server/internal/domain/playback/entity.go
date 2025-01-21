package playback

type (
	Location struct {
		Bucket string
		Key    string
	}
	Locations []Location
	Content   struct {
		ID        string
		Title     string
		Artists   []string
		Album     string
		Genre     string
		Duration  int32
		CoverURL  string
		Locations Locations
	}
	Contents []Content
)
