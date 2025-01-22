package playback

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	domainPlayback "github.com/edwynrrangel/grpc/go/multimedia-client/internal/domain/playback"
)

type handler struct {
	usecase domainPlayback.UseCase
}

func NewHandler(usecase domainPlayback.UseCase) domainPlayback.Handler {
	return &handler{
		usecase: usecase,
	}
}

func (h *handler) Stream(c *gin.Context) {
	reactorValue, exists := c.Get("reactor")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "reactor not found"})
		return
	}
	reactor := reactorValue.(domainPlayback.Reactor)
	eventCh := reactor.Subscribe()

	go func() {
		defer reactor.Close()
		err := h.usecase.Stream(
			c.Request.Context(),
			&domainPlayback.ContentRequest{ContentId: c.Param("content_id")},
			reactor,
		)
		if err != nil {
			reactor.Publish(domainPlayback.StreamEvent{Error: err})
		}
	}()

	c.Stream(func(w io.Writer) bool {
		for event := range eventCh {
			if event.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": event.Error.Error()})
				return false
			}

			if event.Done {

				return false
			}

			_, err := w.Write(event.Data)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return false
			}
		}
		return true
	})
}
