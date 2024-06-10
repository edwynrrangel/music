package resource

import "github.com/gin-gonic/gin"

type Handler interface {
	Public(ctx *gin.Context)
}
