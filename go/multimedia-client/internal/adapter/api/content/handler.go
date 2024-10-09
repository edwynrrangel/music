package content

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/edwynrrangel/grpc/go/multimedia_client/internal/adapter/grpc/content"
	domainContent "github.com/edwynrrangel/grpc/go/multimedia_client/internal/domain/content"
)

type handler struct {
	contentClient content.ContentServiceClient
}

func NewHandler(contentClient content.ContentServiceClient) domainContent.Handler {
	return &handler{
		contentClient: contentClient,
	}
}

func (h *handler) Get(c *gin.Context) {
	got, err := h.contentClient.Get(c.Request.Context(), &content.ContentRequest{Id: c.Param("id")})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, got)
}

func (h *handler) Search(c *gin.Context) {
	got, err := h.contentClient.Search(c.Request.Context(), &content.SearchRequest{Query: c.Query("query")})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, got)
}

func (h *handler) Stream(c *gin.Context) {
	stream, err := h.contentClient.Stream(c.Request.Context(), &content.ContentRequest{Id: c.Param("id")})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Stream(func(w io.Writer) bool {
		for {
			chunk, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					return false
				}
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return false
			}
			_, err = w.Write(chunk.Data)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return false
			}
		}
	})
}
