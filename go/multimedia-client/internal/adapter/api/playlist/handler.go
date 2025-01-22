package playlist

import (
	"net/http"

	domainPlaylist "github.com/edwynrrangel/grpc/go/multimedia-client/internal/domain/playlist"
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

func (h *handler) Create(c *gin.Context) {
	var req domainPlaylist.CreateRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	got, err := h.usecase.Create(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, got)
}

func (h *handler) List(c *gin.Context) {
	got, err := h.usecase.List(c.Request.Context(), &domainPlaylist.ListRequest{
		UserId: c.Param("user_id"),
		Query:  c.Query("query"),
		Page:   int32(c.GetInt("page")),
		Limit:  int32(c.GetInt("limit")),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, got)
}

func (h *handler) Get(c *gin.Context) {
	got, err := h.usecase.Get(c.Request.Context(), &domainPlaylist.PlaylistRequest{
		PlaylistId: c.Param("id"),
		UserId:     c.Param("user_id"),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, got)
}

/* func (h *handler) Manage(c *gin.Context) {
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
} */
