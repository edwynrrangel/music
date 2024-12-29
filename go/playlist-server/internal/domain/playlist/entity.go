package playlist

type (
	CreateRequest struct {
		UserId string
		Name   string
	}
	ListRequest struct {
		UserId string
		Query  string
		Page   int32
		Limit  int32
	}
	Content struct {
		ID    string
		Order int32
	}
	Playlist struct {
		ID     string
		UserId string
		Name   string
		Data   []Content
	}
	Playlists []Playlist
)
