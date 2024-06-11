package resource

import "github.com/gin-gonic/gin"

func RegisterRoutes(app *gin.Engine) {
	h := NewHandler()
	app.Static("/static", "./public/static")
	app.NoRoute(h.Public)
}
