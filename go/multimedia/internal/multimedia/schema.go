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
)
