package content

type (
	SearchResponse struct {
		Data []Content
	}
	StreamResponse struct {
		Data [][]byte
	}
)
