package playlist

type (
	ContentRequest struct {
		ID       string
		Title    string
		Creator  string
		Genre    string
		Duration string
	}
	PlaylistRequest struct {
		ID        string
		UserID    string
		Name      string
		Contents  []ContentRequest
		Operation OperationKey
	}
)
