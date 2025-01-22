package playback

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Stream(c *gin.Context)
}

type UseCase interface {
	Stream(ctx context.Context, data *ContentRequest, reactor Reactor) error
}

type Reactor interface {
	Subscribe() <-chan StreamEvent
	Publish(event StreamEvent)
	Close()
}
