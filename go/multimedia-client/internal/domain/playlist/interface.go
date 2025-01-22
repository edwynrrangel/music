package playlist

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
	// Manage(c *gin.Context)
}

type UseCase interface {
	Create(ctx context.Context, data *CreateRequest) (*Playlist, error)
	List(ctx context.Context, data *ListRequest) (*List, error)
	Get(ctx context.Context, data *PlaylistRequest) (*Playlist, error)
	// AddContent(stream AddContent) error
	// RemoveContent(stream RemoveContent) error
	// PartyMode(ctx context.Context, comm Communication, id string) error
}

type Communication interface {
	WriteJSON(data any) error
	ReadJSON(data any) error
	Close() error
}
