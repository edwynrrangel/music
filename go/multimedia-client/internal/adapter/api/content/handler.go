package content

import (
	"net/http"

	"github.com/gin-gonic/gin"

	domainContent "github.com/edwynrrangel/grpc/go/multimedia-client/internal/domain/content"
)

type handler struct {
	usecase domainContent.UseCase
}

func NewHandler(usecase domainContent.UseCase) domainContent.Handler {
	return &handler{
		usecase: usecase,
	}
}

func (h *handler) Search(c *gin.Context) {
	got, err := h.usecase.Search(c.Request.Context(), &domainContent.SearchRequest{
		Query: c.Query("query"),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, got)
}

func (h *handler) Get(c *gin.Context) {
	got, err := h.usecase.Get(c.Request.Context(), &domainContent.ContentRequest{
		Id: c.Param("id"),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, got)
}
