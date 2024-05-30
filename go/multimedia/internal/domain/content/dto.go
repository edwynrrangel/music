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
	StreamResponse struct {
		Data [][]byte
	}
)
