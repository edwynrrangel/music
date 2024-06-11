package playlist

type (
	PlaylistRequest struct {
		ID        string
		UserID    string
		ContentID string
		Name      string
	}
	RemovePlaylistRequest struct {
		ID     string
		UserID string
	}
)
