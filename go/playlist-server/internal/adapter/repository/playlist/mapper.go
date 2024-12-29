package playlist

import "github.com/edwynrrangel/grpc/go/playlist_server/internal/domain/playlist"

type (
	PlaylistEntity playlist.Playlist
	ContentEntity  playlist.Content
)

func (p *PlaylistEntity) toModel() *Playlist {
	return &Playlist{
		ID:     p.ID,
		UserId: p.UserId,
		Name:   p.Name,
	}
}

func (c *Content) toEntity() *playlist.Content {
	return &playlist.Content{
		ID:    c.ID,
		Order: c.Order,
	}
}

func (p *Playlist) toEntity() *playlist.Playlist {
	contents := make([]playlist.Content, 0, len(p.Contents))
	for _, item := range p.Contents {
		contents = append(contents, *item.toEntity())
	}
	return &playlist.Playlist{
		ID:     p.ID,
		UserId: p.UserId,
		Name:   p.Name,
		Data:   contents,
	}
}

func (p Playlists) toEntity() playlist.Playlists {
	playlists := make(playlist.Playlists, 0)
	for _, playlist := range p {
		playlists = append(playlists, *playlist.toEntity())
	}
	return playlists
}
