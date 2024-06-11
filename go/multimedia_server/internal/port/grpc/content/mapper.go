package content

import "github.com/edwynrrangel/grpc/go/multimedia/internal/domain/content"

// toSearchRequest function converts SearchRequest proto to SearchRequest domain
func (req *SearchRequest) toSearchRequest() *content.SearchRequest {
	return &content.SearchRequest{
		Query: req.Query,
	}
}

// convertToContent function converts Content domain to ContentResponse proto
func convertToContent(content *content.Content) *ContentResponse {
	if content == nil {
		return nil
	}
	return &ContentResponse{
		Id:      content.ID,
		Title:   content.Title,
		Genre:   content.Genre,
		Creator: content.Creator,
	}
}

// convertToSearchResponse function converts SearchResponse domain to SearchResponse proto
func convertToSearchResponse(content *content.SearchResponse) *SearchResponse {
	var list []*ContentResponse
	for _, item := range content.Data {
		list = append(list, convertToContent(&item))
	}
	return &SearchResponse{
		Data: list,
	}
}
