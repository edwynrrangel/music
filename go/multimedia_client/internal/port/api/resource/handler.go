package resource

import (
	"path/filepath"

	"github.com/edwynrrangel/grpc/go/multimedia_client/internal/domain/resource"
	"github.com/gin-gonic/gin"
)

type handler struct{}

func NewHandler() resource.Handler {
	return &handler{}
}

func (h *handler) Public(c *gin.Context) {
	path := filepath.Join("public", "index.html")
	c.File(path)
}
