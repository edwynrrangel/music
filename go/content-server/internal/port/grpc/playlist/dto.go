package playlist

type (
	PlaylistDTO struct {
		ID   string
		Name string
	}
	ListPlaylistsDTO struct {
		UserID string
		Data   []PlaylistDTO
	}
)
