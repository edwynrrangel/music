package content

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
	}
	StreamResponse struct {
		Data [][]byte
	}
)
