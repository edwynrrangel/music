package playlist

import (
	"net/http"

	wsPlaylist "github.com/edwynrrangel/grpc/go/multimedia_client/internal/adapter/websocket/playlist"
	domainPlaylist "github.com/edwynrrangel/grpc/go/multimedia_client/internal/domain/playlist"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type handler struct {
	usecase domainPlaylist.UseCase
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewHandler(usecase domainPlaylist.UseCase) domainPlaylist.Handler {
	return &handler{
		usecase: usecase,
	}
}

func (h *handler) List(c *gin.Context) {
	got, err := h.usecase.List(c.Request.Context(), domainPlaylist.ListFilter{
		UserID: c.Param("user_id"),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, got)
}

func (h *handler) Remove(c *gin.Context) {
	err := h.usecase.Remove(c.Request.Context(), domainPlaylist.RemoveRequest{
		ID:     c.Param("id"),
		UserID: c.Param("user_id"),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *handler) Manage(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()

	err = h.usecase.Manage(c.Request.Context(), wsPlaylist.NewHandler(conn), c.Param("id"))
	if err != nil {
		conn.WriteJSON(gin.H{"error": err.Error()})
		return
	}
}
