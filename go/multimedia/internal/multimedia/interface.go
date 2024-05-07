package multimedia

type UseCase interface {
	SearchContent(request *SearchRequest) (*SearchResponse, error)
}

type Repository interface {
	SearchContent(query string) ([]Content, error)
}
