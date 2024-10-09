package playlist

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	List(c *gin.Context)
	Remove(c *gin.Context)
	Manage(c *gin.Context)
}

type UseCase interface {
	List(ctx context.Context, filters ListFilter) (*ListResponse, error)
	Remove(ctx context.Context, params RemoveRequest) error
	Manage(ctx context.Context, comm Communication, id string) error
}

type Communication interface {
	WriteJSON(data any) error
	ReadJSON(data any) error
	Close() error
}
