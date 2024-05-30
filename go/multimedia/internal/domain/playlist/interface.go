package playlist

import "context"

type UseCase interface {
	Get(context.Context, PlayListRequest) (*PlayList, error)
	Add(context.Context, PlayListRequest) error
	Remove(context.Context, PlayListRequest) error
}

type Repository interface {
	Get(context.Context, string, string) (*PlayList, error)
	Add(context.Context, PlayList) error
	Update(context.Context, string, string, PlayList) error
	Remove(context.Context, string, string) error
}
