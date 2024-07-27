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

type (
	ContentResponse struct {
		ID       string
		Title    string
		Creator  string
		Genre    string
		Duration string
	}
	PlaylistResponse struct {
		ID      string
		UserID  string
		Name    string
		Content []ContentResponse
	}
)
