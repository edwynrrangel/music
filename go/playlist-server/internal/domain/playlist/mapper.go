package playlist

func (c *Content) toDTO() *ContentResponse {
	return &ContentResponse{
		ID:       c.ID,
		Title:    c.Title,
		Creator:  c.Creator,
		Genre:    c.Genre,
		Duration: c.Duration,
	}
}

func (p *Playlist) toDTO() *PlaylistResponse {
	var contentDTO []ContentResponse
	for _, c := range p.Content {
		contentDTO = append(contentDTO, *c.toDTO())
	}
	return &PlaylistResponse{
		ID:      p.ID,
		UserID:  p.UserID,
		Name:    p.Name,
		Content: contentDTO,
	}
}
