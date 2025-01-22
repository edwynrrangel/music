package playlist

import (
	"github.com/edwynrrangel/grpc/go/multimedia-client/internal/domain/playlist"
	"github.com/gorilla/websocket"
)

type handler struct {
	conn *websocket.Conn
}

func NewHandler(conn *websocket.Conn) playlist.Communication {
	return &handler{
		conn: conn,
	}
}

func (w *handler) WriteJSON(data interface{}) error {
	return w.conn.WriteJSON(data)
}

func (w *handler) ReadJSON(data interface{}) error {
	return w.conn.ReadJSON(data)
}

func (w *handler) Close() error {
	return w.conn.Close()
}
