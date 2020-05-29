package ws

import (
	"github.com/gorilla/websocket"
	"time"
)

func readRoutine(conn *websocket.Conn) {
	for {
		select {
		default:
			typ, body, err := conn.ReadMessage()
			if err != nil {
				logger.Warn("read message err, closed:", err)
			} else {
				logger.Debug("message:", typ, string(body))
			}
		}
	}
}

func write(conn *websocket.Conn, msgType int, message []byte) error {
	if err := conn.SetWriteDeadline(time.Now().Add(time.Second * 3)); err != nil {
		logger.Warn("set write dead line err:", err)
		return err
	}
	if err := conn.WriteMessage(msgType, message); err != nil {
		logger.Error("write message err:", err)
		return err
	}
	return nil
}
