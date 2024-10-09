package playlist

type (
	ListFilter struct {
		UserID string
	}
	RemoveRequest struct {
		ID     string
		UserID string
	}
	DataList struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	ListResponse struct {
		Data []DataList `json:"data"`
	}
)
