package content

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Search(c *gin.Context)
	Get(c *gin.Context)
}

type UseCase interface {
	Search(ctx context.Context, data *SearchRequest) (*ContentList, error)
	Get(ctx context.Context, data *ContentRequest) (*Content, error)
}
