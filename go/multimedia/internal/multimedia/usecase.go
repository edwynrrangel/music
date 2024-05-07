package multimedia

type usecase struct {
	repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &usecase{repo: repo}
}

func (u *usecase) SearchContent(request *SearchRequest) (*SearchResponse, error) {
	contents, err := u.repo.SearchContent(request.Query)
	if err != nil {
		return nil, err
	}

	return &SearchResponse{
		Data: contents,
	}, nil
}
