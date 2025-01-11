package playlist

type (
	CreateRequest struct {
		UserId string
		Name   string
		Mode   string
	}
	ListRequest struct {
		UserId string
		Query  string
		Page   int32
		Limit  int32
	}
	AddContentEvent struct {
		ContentId string
		Order     int32
	}
	SkipContentEvent struct {
		ForceSkip bool
	}
	Action struct {
		AddContent  *AddContentEvent
		SkipContent *SkipContentEvent
	}
	PartyModeRequest struct {
		PlaylistId string
		Action     Action
	}
	Content struct {
		ID    string
		Order int32
	}
	Playlist struct {
		ID     string
		UserId string
		Name   string
		Mode   string
		Data   []Content
	}
	Playlists []Playlist
)
