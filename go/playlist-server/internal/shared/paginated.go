package shared

type Paginated[T any] struct {
	Page  int32
	Limit int32
	Total int64
	Data  []T
}
