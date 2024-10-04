package content

import "github.com/edwynrrangel/grpc/go/multimedia_server/internal/domain/content"

// convertToContent function converts Content domain to ContentResponse proto
func convertToContent(content *content.Content) *ContentResponse {
	if content == nil {
		return nil
	}
	return &ContentResponse{
		Id:       content.ID,
		Title:    content.Title,
		Genre:    content.Genre,
		Creator:  content.Creator,
		Duration: content.Duration,
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
