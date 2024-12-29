package playlist

type (
	CreateRequest struct {
		UserID string
		Name   string
	}
	ListRequest struct {
		UserID string
		Query  string
		Page   int32
		Limit  int32
	}
	Playlist struct {
		ID     string
		UserID string
		Name   string
		Data   []string
	}
	Playlists []Playlist
)
