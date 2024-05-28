package multimedia

type (
	SearchRequest struct {
		Query string
	}
)

type (
	SearchResponse struct {
		Data []Content
	}
	ContentResponse struct {
		ID      string
		Title   string
		Genre   string
		Creator string
		Url     string
	}
	StreamResponse struct {
		Data [][]byte
	}
)
