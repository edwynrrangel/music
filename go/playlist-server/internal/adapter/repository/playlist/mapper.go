package playlist

import "github.com/edwynrrangel/grpc/go/playlist_server/internal/domain/playlist"

type (
	PlaylistEntity playlist.Playlist
)

func (p *PlaylistEntity) toModel() *Playlist {
	return &Playlist{
		ID:     p.ID,
		UserID: p.UserID,
		Name:   p.Name,
	}
}

func (p *Playlist) toEntity() *playlist.Playlist {
	return &playlist.Playlist{
		ID:     p.ID,
		UserID: p.UserID,
		Name:   p.Name,
		Data:   p.Contents,
	}
}

func (p Playlists) toEntity() playlist.Playlists {
	playlists := make(playlist.Playlists, 0)
	for _, playlist := range p {
		playlists = append(playlists, *playlist.toEntity())
	}
	return playlists
}

/*
func (p *PlaylistRequest) ToEntity() *Playlist {
	contents := make([]Content, 0)
	for _, content := range p.Contents {
		contents = append(contents, *content.ToEntity())
	}
	return &Playlist{
		ID:       p.ID,
		UserID:   p.UserID,
		Name:     p.Name,
		Contents: contents,
	}
} */
