package playlist

func (c *ContentRequest) ToEntity() *Content {
	return &Content{
		ID:       c.ID,
		Title:    c.Title,
		Creator:  c.Creator,
		Genre:    c.Genre,
		Duration: c.Duration,
	}
}

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
}
